package main

import (
	"os"
	"fmt"
	"bufio"
	"time"
	"flag"
	"regexp"
	"strconv"
	"strings"
	"path/filepath"
	"io/ioutil"
	"github.com/Unknwon/goconfig"
	"www.baidu.com/golang-lib/log"
)

//获取gsid那行日志当前对应的日志文件名称
func getCurLogName(appName string, logName string, logTime string) (logFile string, err error) {
	confDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Logger.Critical("get conf file error, %s", err)
		return "", fmt.Errorf("get conf file error, %s", err)
	}
	conf_file := confDir + "/conf/" + appName + ".ini"

	cfg, err := goconfig.LoadConfigFile(conf_file)
	if err != nil {
		log.Logger.Critical("load conf error %s", err)
	}

	log.Logger.Info(appName, logName, logTime)
	rotateType, _ := cfg.GetValue(logName, "rotate_type")
	logFormat, _ := cfg.GetValue(logName, "log_format")
	bakPath, _ := cfg.GetValue(logName, "bak_path")
	bakName, _ := cfg.GetValue(logName, "bak_name")
	bakLogFormat, _ := cfg.GetValue(logName, "bak_format")
	if rotateType == "0" {
		return getCurLogNameRegular(bakLogFormat, logFormat, logTime)
	} else if rotateType == "1" {
		return getCurLogNameIrregular(bakPath, bakName, logFormat, logTime)
	} else {
		log.Logger.Critical("log rotate type not supported, %s", rotateType)
		return "", fmt.Errorf("rotate type invalid, %s", rotateType)
	}
}

//对于按时间定期切割的日志，切分规则固定，归类为规则切分
func getCurLogNameRegular(bakLogFormat string, logFormat string, logTime string) (logName string, err error) {
	//通过比较文件名判断日志是否已经切分
	bakLogFile := getLogByFormat(bakLogFormat, logTime)
	if isExist(bakLogFile) {
		return bakLogFile, nil
	} else {
		logFile := getLogByFormat(logFormat, logTime)
		return logFile, nil
	}
}

//对于按大小、或按时间但不固定切分的日志，归类为不规则切分
func getCurLogNameIrregular(bakPath string, bakName string, logFormat string, logTime string) (logName string, err error) {
	//从日志备份目录下找出符合日志文件备份格式的文件，通过mtime与日志采集时间对比来确定当前文件名
	//mtime > 采集时间 的文件里，mtime最小的那个文件就是该行日志切分后所在的文件
	//如果不存在这样的文件，则没有切分
	logTimeStamp := timeStr2Stamp(strings.Replace(logTime, "_", " ", -1))
	mtimeMin := int64(0)
	curLogName := ""
    fileList, err := ioutil.ReadDir(bakPath)
    if err != nil {
		log.Logger.Critical("read baklog dir error %s", err)
		return "", fmt.Errorf("read baklog dir error, %s, %s", err, bakPath)
    }

    for _, v := range fileList {
		if v.IsDir() {		//过滤目录
			continue
		}
		matched, err := regexp.MatchString(bakName, v.Name()) 
		if !matched || err != nil {		//过滤不匹配日志名的文件
			continue
		}

		logFile := bakPath + "/" + v.Name()
		mtime, err := getMtime(logFile)
		if err != nil || mtime < logTimeStamp {
			continue
		}
		if mtimeMin == 0 || mtime < mtimeMin {
			mtimeMin = mtime
			curLogName = v.Name()
		}
    }
	if curLogName != "" {
		return bakPath + "/" + curLogName, nil
	}

	//在备份文件中没有找到符合条件的，则说明没有切分，返回当前正在打印的日志文件
	return logFormat, nil
}

//时间字符串转时间戳
func timeStr2Stamp(timeStr string) (timestamp int64) {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, timeStr, loc)
	return theTime.Unix()
}

//判断文件是否存在
func isExist(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	return false
}

//根据配置生成日志名
func getLogByFormat(logFormat string, logTime string) (log string) {
	reg_year := regexp.MustCompile(`\{year\}`)
	reg_month := regexp.MustCompile(`\{month\}`)
	reg_day := regexp.MustCompile(`\{day\}`)
	reg_hour := regexp.MustCompile(`\{hour\}`)
	reg_minute := regexp.MustCompile(`\{minute\}`)

	year := []byte(logTime)[:4]
	month := []byte(logTime)[5:7]
	day := []byte(logTime)[8:10]
	hour := []byte(logTime)[11:13]
	minute := []byte(logTime)[14:16]

	log = logFormat
	log = string(reg_year.ReplaceAll([]byte(log), year))
	log = string(reg_month.ReplaceAll([]byte(log), month))
	log = string(reg_day.ReplaceAll([]byte(log), day))
	log = string(reg_hour.ReplaceAll([]byte(log), hour))
	log = string(reg_minute.ReplaceAll([]byte(log), minute))

	//fmt.Println(log)
	return log
}

//对于按大小切割的日志，根据文件最近修改时间判断所在的日志文件名
func getMtime(file string) (mtime int64, err error) {
	f, err := os.Open(file)
    if err != nil {
		return int64(0), fmt.Errorf("open file error, %s, %s", err, file)
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return int64(0), fmt.Errorf("stat file error, %s, %s", err, file)
	}
	return fi.ModTime().Unix(), nil
}

//到指定文件的指定位移，读取一行
func readLineByOffset(file string, offset int64) (line string, err error) {
	fh, err := os.Open(file)
	if err != nil {
		log.Logger.Critical("open file failed, %s", err)
		return "", err
	}
	_, err = fh.Seek(offset, 0)
	if err != nil {
		log.Logger.Critical("seek file failed, %s", err)
		return "", err
	}
	reader := bufio.NewReader(fh)
	line, err = reader.ReadString('\n')
	if err != nil {
		log.Logger.Critical("read file failed, %s", err)
		return "", err
	}
	return line, nil
}

func main() {
	log.Init("trace_gsid", "INFO", "./local_data/logs", false, "H", 5)
	logTime := flag.String("time", "", "log time")
	appName := flag.String("app", "", "app name")
	logName := flag.String("log", "", "log name in conf file")
	offset := flag.String("offset", "", "log offset")
	flag.Parse()

	log.Logger.Info(*logTime, *appName, *logName, *offset)

	logFile, err := getCurLogName(*appName, *logName, *logTime)
	log.Logger.Info("get log over, %s", logFile)
	if err != nil {
		log.Logger.Critical("get cur log info failed, %s", err)
		return
	}

	offsetInt64, err := strconv.ParseInt(*offset, 10, 64)
	if err != nil {
		log.Logger.Critical("parse offset error, %s", err)
		return
	}
	logContent, err := readLineByOffset(logFile, offsetInt64)
	if err != nil {
		log.Logger.Critical("read log by offset failed, %s", err)
		return
	}

	fmt.Println(logContent)
}
