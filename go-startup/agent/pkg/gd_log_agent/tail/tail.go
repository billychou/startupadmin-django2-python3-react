package tail

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"

	"gd_log_agent/tail/logrotate"
	"gd_log_agent/tail/ratelimiter"
	"gd_log_agent/tail/util"
	"gd_log_agent/tail/watch"
	"gopkg.in/tomb.v1"
	//"gd_log_agent/tail/vendor/gopkg.in/tomb.v1"
)

var (
	ErrStop = errors.New("tail should now stop")
)

type Line struct {
	Text    string
	Offset  int64  //add by mushi
	Time    string //format  2018-04-01_00:00:00
	LogFile string //log path and log name
	//Err  error // Error from tail
}

// NewLine returns a Line with present time.
//func NewLine(text string) *Line {
//	return &Line{text, time.Now(), nil, 0}
//}

// SeekInfo represents arguments to `os.Seek`
type SeekInfo struct {
	Offset int64
	Whence int // os.SEEK_*
}

type logger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

// Config is used to specify how a file must be tailed.
type Config struct {
	// File-specifc
	Location    *SeekInfo // Seek to this location before tailing
	ReOpen      bool      // Reopen recreated files (tail -F)
	MustExist   bool      // Fail early if the file does not exist
	Poll        bool      // Poll for file changes instead of using inotify
	Pipe        bool      // Is a named pipe (mkfifo)
	RateLimiter *ratelimiter.LeakyBucket

	// Generic IO
	Follow      bool // Continue looking for new lines (tail -f)
	MaxLineSize int  // If non-zero, split longer lines into multiple lines

	// Logger, when nil, is set to tail.DefaultLogger
	// To disable logging: set field to tail.DiscardingLogger
	Logger logger

	//日志个性化字段
	GsidReg string
}

//日志切分配置，初次启动时加载文本配置内容，并初始化
type LogRotateConfig struct {
	LogFormat     string //配置中的日志格式字段：log_format, 包含日志路径和日志名字
	LogRotateType string //日志切分类型，0：按时间、1：按大小
	//LogRotateGap	int		//日志切分周期，单位为小时
	//LastRotateTime	string		//上次日志切分时间
}

type Tail struct {
	Filename string
	Lines    chan *Line
	Config
	LogRotateConfig
	FWChan chan int //用于在日志切分后使老的watcher线程退出

	file   *os.File
	reader *bufio.Reader

	watcher watch.FileWatcher
	changes *watch.FileChanges

	tomb.Tomb // provides: Done, Kill, Dying

	lk sync.Mutex
}

var (
	// DefaultLogger is used when Config.Logger == nil
	DefaultLogger = log.New(os.Stderr, "", log.LstdFlags)
	// DiscardingLogger can be used to disable logging output
	DiscardingLogger = log.New(ioutil.Discard, "", 0)
)

// TailFile begins tailing the file. Output stream is made available
// via the `Tail.Lines` channel. To handle errors during tailing,
// invoke the `Wait` or `Err` method after finishing reading from the
// `Lines` channel.
func TailFile(filename string, config Config, logRotateConfig LogRotateConfig) (*Tail, error) {
	if config.ReOpen && !config.Follow {
		util.Fatal("cannot set ReOpen without Follow.")
	}

	t := &Tail{
		Filename:        filename,
		Lines:           make(chan *Line),
		Config:          config,
		LogRotateConfig: logRotateConfig,
		FWChan:          make(chan int),
	}

	// when Logger was not specified in config, use default logger
	if t.Logger == nil {
		t.Logger = log.New(os.Stderr, "", log.LstdFlags)
	}

	if t.Poll {
		t.watcher = watch.NewPollingFileWatcher(filename, t.LogRotateType, t.LogFormat)
		//非poll模式暂时忽略
	} else {
		//t.watcher = watch.NewInotifyFileWatcher(filename)
		return nil, nil
	}

	if t.MustExist {
		var err error
		t.file, err = OpenFile(t.Filename)
		if err != nil {
			return nil, err
		}
	}

	go t.tailFileSync()

	return t, nil
}

// Return the file's current position, like stdio's ftell().
// But this value is not very accurate.
// it may readed one line in the chan(tail.Lines),
// so it may lost one line.
func (tail *Tail) Tell() (offset int64, err error) {
	if tail.file == nil {
		tail.Logger.Printf("tell error 1 %s", err)
		return
	}
	offset, err = tail.file.Seek(0, os.SEEK_CUR)
	if err != nil {
		tail.Logger.Printf("tell error 2 %s", err)
		return
	}

	tail.lk.Lock()
	defer tail.lk.Unlock()
	if tail.reader == nil {
		tail.Logger.Printf("tell error 3 %s", err)
		return
	}

	offset -= int64(tail.reader.Buffered())
	return
}

// Stop stops the tailing activity.
func (tail *Tail) Stop() error {
	tail.Kill(nil)
	tail.Logger.Printf("stop ...")
	return tail.Wait()
}

// StopAtEOF stops tailing as soon as the end of the file is reached.
func (tail *Tail) StopAtEOF() error {
	tail.Kill(errStopAtEOF)
	tail.Logger.Printf("stop at eof")
	return tail.Wait()
}

var errStopAtEOF = errors.New("tail: stop at eof")

func (tail *Tail) close() {
	close(tail.Lines)
	tail.closeFile()
}

func (tail *Tail) closeFile() {
	if tail.file != nil {
		tail.file.Close()
		tail.file = nil
	}
}

func (tail *Tail) reopen() error {
	tail.closeFile()
	for {
		var err error
		tail.file, err = OpenFile(tail.Filename)
		if err != nil {
			if os.IsNotExist(err) {
				tail.Logger.Printf("Waiting for %s to appear...", tail.Filename)
				if err := tail.watcher.BlockUntilExists(&tail.Tomb); err != nil {
					//tail.Logger.Printf("xxxxxxxxxxxxx")
					if err == tomb.ErrDying {
						tail.Logger.Printf("reopen error %s", err)
						return err
					}
					tail.Logger.Printf("detect fail %s", err)
					return fmt.Errorf("Failed to detect creation of %s: %s", tail.Filename, err)
				}
				//tail.Logger.Printf("1111111111")
				continue
			}
			tail.Logger.Printf("open fail %s", err)
			return fmt.Errorf("Unable to open file %s: %s", tail.Filename, err)
		}
		break
	}
	return nil
}

func (tail *Tail) readLine() (string, error) {
	tail.lk.Lock()
	line, err := tail.reader.ReadString('\n')
	tail.lk.Unlock()
	if err != nil {
		// Note ReadString "returns the data read before the error" in
		// case of an error, including EOF, so we return it as is. The
		// caller is expected to process it if err is EOF.
		return line, err
	}

	line = strings.TrimRight(line, "\n")

	return line, err
}

func (tail *Tail) tailFileSync() {
	defer tail.Done()
	defer tail.close()

	if !tail.MustExist {
		// deferred first open.
		err := tail.reopen()
		if err != nil {
			if err != tomb.ErrDying {
				tail.Logger.Printf("sync tail return error %s", err)
				tail.Kill(err)
			}
			return
		}
	}

	// Seek to requested location on first open of the file.
	if tail.Location != nil {
		_, err := tail.file.Seek(tail.Location.Offset, tail.Location.Whence)
		tail.Logger.Printf("Seeked %s - %+v\n", tail.Filename, tail.Location)
		if err != nil {
			tail.Killf("Seek error on %s: %s", tail.Filename, err)
			tail.Logger.Printf("seek error %s", err)
			return
		}
	}

	tail.openReader()

	var offset int64
	var err error

	//add by mushi reg
	reg := regexp.MustCompile(tail.GsidReg)

	// Read line by line.
	for {
		// do not seek in named pipes
		if !tail.Pipe {
			// grab the position in case we need to back up in the event of a half-line
			offset, err = tail.Tell()
			if err != nil {
				tail.Kill(err)
				tail.Logger.Printf("tell return error %s", err)
				return
			}
		}

		line, err := tail.readLine()

		// Process `line` even if err is EOF.
		if err == nil {
			// add by mushi, gsid
			gsid := ""
			subMatch := reg.FindStringSubmatch(line)
			if len(subMatch) >= 2 {
				gsid = subMatch[1]
			} else {
				//gsid = "xxxx"
				continue
			}

			//2018-01-01_00:00:00
			now := string([]byte(time.Now().String())[:19])
			nowFormated := strings.Replace(now, " ", "_", -1)

			cooloff := !tail.sendLine(gsid, offset, nowFormated, tail.Filename) // add by mushi , offset
			if cooloff {
				// Wait a second before seeking till the end of
				// file when rate limit is reached.
				msg := ("Too much log activity; waiting a second " +
					"before resuming tailing")
				//add offset and modify by mushi
				tail.Logger.Printf(msg)
				//tail.Lines <- &Line{msg, time.Now(), errors.New(msg), offset}
				//tail.Lines <- &Line{msg, offset}
				select {
				case <-time.After(time.Second):
				case <-tail.Dying():
					tail.Logger.Printf("dying %s", err)
					return
				}
				if err := tail.seekEnd(); err != nil {
					tail.Kill(err)
					tail.Logger.Printf("seek end error %s", err)
					return
				}
			}
		} else if err == io.EOF {
			if !tail.Follow {
				//trace场景中Follow都为True，这种case暂时忽略
				//				if line != "" {
				//					tail.sendLine(line, offset) // add by mushi , offset
				//				}
				return
			}

			if tail.Follow && line != "" {
				// this has the potential to never return the last line if
				// it's not followed by a newline; seems a fair trade here
				err := tail.seekTo(SeekInfo{Offset: offset, Whence: 0})
				if err != nil {
					tail.Kill(err)
					tail.Logger.Printf("seekto error %s", err)
					return
				}
			}

			// When EOF is reached, wait for more data to become
			// available. Wait strategy is based on the `tail.watcher`
			// implementation (inotify or polling).
			err := tail.waitForChanges()
			if err != nil {
				tail.Logger.Printf("wait change %s", err)
				if err != ErrStop {
					tail.Kill(err)
				}
				return
			}
		} else {
			// non-EOF error
			tail.Logger.Printf("raading error %s", err)
			tail.Killf("Error reading %s: %s", tail.Filename, err)
			return
		}

		select {
		case <-tail.Dying():
			if tail.Err() == errStopAtEOF {
				continue
			}
			tail.Logger.Printf("tail dying %s", err)
			return
		default:
		}
	}
}

// waitForChanges waits until the file has been appended, deleted,
// moved or truncated. When moved or deleted - the file will be
// reopened if ReOpen is true. Truncated files are always reopened.
func (tail *Tail) waitForChanges() error {
	if tail.changes == nil {
		//tail.Logger.Printf("wait change 0")
		pos, err := tail.file.Seek(0, os.SEEK_CUR)
		if err != nil {
			tail.Logger.Printf("wait change error 1 %s", err)
			return err
		}
		//tail.changes, err = tail.watcher.ChangeEvents(&tail.Tomb, pos)
		tail.changes, err = tail.watcher.ChangeEvents(&tail.Tomb, pos, tail.FWChan)
		if err != nil {
			tail.Logger.Printf("wait for change return error %s", err)
			return err
		}
		//tail.Logger.Printf("new changes created")
	}

	select {
	case <-tail.changes.Modified:
		return nil
	case <-tail.changes.Deleted:
		tail.changes = nil
		if tail.ReOpen {
			// XXX: we must not log from a library.
			tail.Logger.Printf("Re-opening moved/deleted file %s ...", tail.Filename)
			if err := tail.reopen(); err != nil {
				tail.Logger.Printf("wait change error 2 %s", err)
				return err
			}
			tail.Logger.Printf("Successfully reopened %s", tail.Filename)
			tail.openReader()
			return nil
		} else {
			tail.Logger.Printf("Stopping tail as file no longer exists: %s", tail.Filename)
			return ErrStop
		}
	case <-tail.changes.Truncated:
		// Always reopen truncated files (Follow is true)
		tail.Logger.Printf("Re-opening truncated file %s ...", tail.Filename)
		if err := tail.reopen(); err != nil {
			tail.Logger.Printf("wait change error 3 %s", err)
			return err
		}
		tail.Logger.Printf("Successfully reopened truncated %s", tail.Filename)
		tail.openReader()
		return nil
	//为了处理日志切分的情况
	//如果日志既没有删除、改变、截取，那可能是日志切分后打印到新文件中了,需要判断一下
	//如果是，则更新日志文件名，重新打开，如果否，则继续
	case <-tail.changes.Rotated:
		//强制让老的changes进程退出，下次轮询会重新生成一个changes进程
		tail.changes = nil
		tail.FWChan <- 0 //ChangeEvent里监听FWChan管道，有数据则退出
		if err := tail.reopenRotatedLog(); err != nil {
			tail.Logger.Printf("deal rotate error %s", err)
			return err
		}
		tail.Logger.Printf("Successfully reopened rotated log %s", tail.Filename)
		return nil
	case <-tail.Dying():
		tail.Logger.Printf("wait change error 4")
		return ErrStop
	}
	panic("unreachable")
}

//add by mushi
//处理日志切分后，重新打开文件的逻辑
func (tail *Tail) reopenRotatedLog() error {
	tail.Logger.Printf("reopened rotated log called")
	//关闭老文件
	tail.closeFile()
	//修改日志文件名,同时需要设置tail结构体中的Filename和watcher两个字段的内容
	//tail.Filename 是读取的文件指针打开的文件
	newLog, err := logrotate.GetCurLog(string(tail.LogRotateType), tail.LogFormat)
	if err != nil {
		tail.Logger.Printf("get new log error, %s", err)
		return fmt.Errorf("get new log error, %s", err)
	}
	s := reflect.ValueOf(tail).Elem()
	s.Field(0).SetString(newLog)
	//tail.watcher 是监测文件内容变化的监视器，里面也有一个Filename字段，标识被监测的文件
	tail.watcher = watch.NewPollingFileWatcher(newLog, tail.LogRotateType, tail.LogFormat)
	//tail.Logger.Printf("watcher update %+v\n", tail.watcher)
	//重新打开
	err = tail.reopen()
	//文件重新打开后，需要重新生成一次reader，否则还会从老的bufio里读取
	tail.openReader()
	if err != nil {
		tail.Logger.Printf("reopen rotated log error, %s", err)
		return fmt.Errorf("reopen rotated log error, %s", err)
	}
	return nil
}

func (tail *Tail) openReader() {
	if tail.MaxLineSize > 0 {
		// add 2 to account for newline characters
		tail.reader = bufio.NewReaderSize(tail.file, tail.MaxLineSize+2)
	} else {
		tail.reader = bufio.NewReader(tail.file)
	}
}

func (tail *Tail) seekEnd() error {
	return tail.seekTo(SeekInfo{Offset: 0, Whence: os.SEEK_END})
}

func (tail *Tail) seekTo(pos SeekInfo) error {
	_, err := tail.file.Seek(pos.Offset, pos.Whence)
	if err != nil {
		tail.Logger.Printf("seek err %s", err)
		return fmt.Errorf("Seek error on %s: %s", tail.Filename, err)
	}
	// Reset the read buffer whenever the file is re-seek'ed
	tail.reader.Reset(tail.file)
	return nil
}

// sendLine sends the line(s) to Lines channel, splitting longer lines
// if necessary. Return false if rate limit is reached.
//add by mushi , offset
func (tail *Tail) sendLine(line string, offset int64, now string, logFile string) bool {
	lines := []string{line}

	// Split longer lines
	if tail.MaxLineSize > 0 && len(line) > tail.MaxLineSize {
		lines = util.PartitionString(line, tail.MaxLineSize)
	}

	for _, line := range lines {
		tail.Lines <- &Line{line, offset, now, logFile} //add offset by mushi
	}

	if tail.Config.RateLimiter != nil {
		ok := tail.Config.RateLimiter.Pour(uint16(len(lines)))
		if !ok {
			tail.Logger.Printf("Leaky bucket full (%v); entering 1s cooloff period.\n",
				tail.Filename)
			return false
		}
	}

	return true
}

// Cleanup removes inotify watches added by the tail package. This function is
// meant to be invoked from a process's exit handler. Linux kernel may not
// automatically remove inotify watches after the process exits.
func (tail *Tail) Cleanup() {
	watch.Cleanup(tail.Filename)
}
