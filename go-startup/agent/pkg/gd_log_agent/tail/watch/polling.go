package watch

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"gd_log_agent/tail/logrotate"
	"gd_log_agent/tail/util"
	"gopkg.in/tomb.v1"
)

// PollingFileWatcher polls the file for changes.
type PollingFileWatcher struct {
	Filename      string
	Size          int64
	LogRotateType string
	LogFormat     string
}

func NewPollingFileWatcher(filename string, logRotateType string, logFormat string) *PollingFileWatcher {
	fw := &PollingFileWatcher{filename, 0, logRotateType, logFormat}
	return fw
}

var POLL_DURATION time.Duration

func (fw *PollingFileWatcher) BlockUntilExists(t *tomb.Tomb) error {
	for {
		if _, err := os.Stat(fw.Filename); err == nil {
			return nil
		} else if !os.IsNotExist(err) {
			return err
		}
		select {
		case <-time.After(POLL_DURATION):
			continue
		case <-t.Dying():
			return tomb.ErrDying
		}
	}
	panic("unreachable")
}

//func (fw *PollingFileWatcher) ChangeEvents(t *tomb.Tomb, pos int64) (*FileChanges, error) {
func (fw *PollingFileWatcher) ChangeEvents(t *tomb.Tomb, pos int64, FWChan chan int) (*FileChanges, error) {
	origFi, err := os.Stat(fw.Filename)
	if err != nil {
		return nil, err
	}

	changes := NewFileChanges()
	var prevModTime time.Time

	// XXX: use tomb.Tomb to cleanly manage these goroutines. replace
	// the fatal (below) with tomb's Kill.

	fw.Size = pos

	go func() {
		prevSize := fw.Size
		for {
			select {
			case <-t.Dying():
				return
			case c := <-FWChan:
				fmt.Println("receive exit cmd, %s", string(c))
				return
			default:
			}

			time.Sleep(POLL_DURATION)
			fi, err := os.Stat(fw.Filename)
			if err != nil {
				// Windows cannot delete a file if a handle is still open (tail keeps one open)
				// so it gives access denied to anything trying to read it until all handles are released.
				if os.IsNotExist(err) || (runtime.GOOS == "windows" && os.IsPermission(err)) {
					// File does not exist (has been deleted).
					changes.NotifyDeleted()
					return
				}

				// XXX: report this error back to the user
				util.Fatal("Failed to stat file %v: %v", fw.Filename, err)
			}

			// File got moved/renamed?
			if !os.SameFile(origFi, fi) {
				changes.NotifyDeleted()
				return
			}

			// File got truncated?
			fw.Size = fi.Size()
			if prevSize > 0 && prevSize > fw.Size {
				changes.NotifyTruncated()
				prevSize = fw.Size
				continue
			}
			// File got bigger?
			if prevSize > 0 && prevSize < fw.Size {
				changes.NotifyModified()
				prevSize = fw.Size
				continue
			}
			prevSize = fw.Size

			// File was appended to (changed)?
			modTime := fi.ModTime()
			if modTime != prevModTime {
				prevModTime = modTime
				changes.NotifyModified()
				continue // add by mushi
			}

			//add by mushi
			//File was rotated? (need to reopen new log file)
			rotated := logrotate.IsRotated(fw.LogRotateType, fw.LogFormat, fw.Filename)
			if rotated {
				//fmt.Printf("%+v\n", fw)
				changes.NotifyRotated()
			}
		}
	}()

	return changes, nil
}

func init() {
	POLL_DURATION = 250 * time.Millisecond
}
