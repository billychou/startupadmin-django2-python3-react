//deal with log rotate rules

package logrotate

import (
	"time"
	"fmt"
	"regexp"
)

//func GetLogByTime()

func GetCurLog(logRotateType string, logFormat string) (log string, err error) {
	if logRotateType == "0" {
		newLog := GetCurLogByTime(logFormat)
		return newLog, nil
	} else if logRotateType == "1" {
		return "", fmt.Errorf("log conf rotate type not supported, type=[%s]", logRotateType)
	} else {
		return "", fmt.Errorf("log conf rotate type error, type=[%s]", logRotateType)
	} }

//判断日志是否进行了切分
func IsRotated(logRotateType string, logFormat string, oldLog string) bool {
	//fmt.Println("is rotated called")
	//按时间切分的日志
	if logRotateType == "0" {
		curLog := GetCurLogByTime(logFormat)
		if curLog != oldLog {
			fmt.Println("not equal, %s, %s", oldLog, curLog)
			return true
		} else {
			//fmt.Println("isequal, %s, %s", oldLog, curLog)
			return false
		}
	//按照size切分, 这种情况都是正在打印的日志名不变，mv成备份名切割的，truncated可以覆盖
	//所以这里先都返回false
	} else if logRotateType == "1" {
		return false
	} else {
		return false
	}
}

//根据日志切分和命名规则，结合当前时间生成理论上的当前日志名
func GetCurLogByTime(log_format string) (log string) {
	reg_year := regexp.MustCompile(`\{year\}`)
	reg_month := regexp.MustCompile(`\{month\}`)
	reg_day := regexp.MustCompile(`\{day\}`)
	reg_hour := regexp.MustCompile(`\{hour\}`)
	reg_minute := regexp.MustCompile(`\{minute\}`)

	now := time.Now().String()
	year := []byte(now)[:4]
	month := []byte(now)[5:7]
	day := []byte(now)[8:10]
	hour := []byte(now)[11:13]
	minute := []byte(now)[14:16]

	log = log_format
	log = string(reg_year.ReplaceAll([]byte(log), year))
	log = string(reg_month.ReplaceAll([]byte(log), month))
	log = string(reg_day.ReplaceAll([]byte(log), day))
	log = string(reg_hour.ReplaceAll([]byte(log), hour))
	log = string(reg_minute.ReplaceAll([]byte(log), minute))

	//fmt.Println(log)
	return log
}

