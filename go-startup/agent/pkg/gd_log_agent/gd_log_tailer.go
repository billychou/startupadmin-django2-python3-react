package main

import (
	"os"
	"fmt"
	"time"
	"io/ioutil"
	"os/exec"
	"strings"
	"runtime"
	"regexp"
	//"reflect"
	"github.com/Unknwon/goconfig"
	"www.baidu.com/golang-lib/log"
	"gd_log_agent/tail"
	//"gd_log_agent/logconf"
)

func runCmd(cmd string) (ret string, err error) {
	c := exec.Command("sh", "-c", cmd)
	stdout, err := c.StdoutPipe()
	if err != nil {
		log.Logger.Critical(err)
		return "", err
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := c.Start(); err != nil {
		log.Logger.Critical(err)
		return "", err
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	return string(opBytes), err
}

//获取本机ip
func getLocalIp() (ip string) {
	ret, err := runCmd("hostname -i")
	if err != nil {
		log.Logger.Critical(err)
		return
	}
	ip = strings.Replace(string(ret), "\n", "", -1)
	//log.Println(ip)
	return ip
}
//根据机器ip，获取机器所属app
func getAppByIp(ip string) (app string) {
	cmd := "armory -ei " + ip +" | grep product_name | head -n1 | awk '{print $2}'"
	ret, err := runCmd(cmd)
	if err != nil {
		log.Logger.Critical(err)
		return
	}
	app = strings.Replace(string(ret), "\n", "", -1)
	//log.Println(ip)
	return app
}

//tail日志内容
func tailLog(ip string, logAlias string, filename string, config tail.Config, 
			logRotateConfig tail.LogRotateConfig, app string) {
	tails, err := tail.TailFile(filename, config, logRotateConfig)
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}
	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		//fmt.Println(msg.Time, msg.Text, ip, app, msg.LogFile, msg.Offset)
		log.Logger.Info("%s,%s|%s|%s|%s|%d", msg.Text, msg.Time, ip, app, logAlias, msg.Offset)
		//log.Logger.Info("%s %s %s %s %s %d", msg.Time, msg.Text, ip, app, logAlias, msg.Offset)
		//log.Logger.Info("%s %s %s %s %s %d", msg.Time, msg.Text, ip, app, msg.LogFile, msg.Offset)
	}
}


//根据配置生成日志名
func getLog(log_format string) (log string) {
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

	fmt.Println(log)
	return log
}

//判断文件是否存在
func isExist(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	return false
}

func main() {
	//启动先判断配置文件是否存在，不存在的直接退出
	ip := getLocalIp()
	app := getAppByIp(ip)
	conf_file := "conf/" + app + ".ini"
	if !isExist(conf_file) {
		return
	}

	log.Init("trace_index", "INFO", "./local_data/logs", false, "H", 5)

	var log_chan chan int = make(chan int)
	runtime.GOMAXPROCS(2)
	//for test
	//conf_file = "conf/erlangshen-aos-amaps.ini"

	//logconf.InitConf()
	//fmt.Println("-----------------------")
	//s := "AmapAosAmapsConf"
	//f := reflect.ValueOf("A")
	//fmt.Println(f)
	//fmt.Println(logconf.AmapAosAmapsConf[0].GsidReg)
	//fmt.Println(logconf.f[0].GsidReg)
	//fmt.Println("-----------------------")
	cfg, err := goconfig.LoadConfigFile(conf_file)
	if err != nil {
		log.Logger.Critical("load conf error %s", err)
	}

	sections := cfg.GetSectionList()
	fmt.Println(sections)

	for i := 0; i < len(sections); i++ {
		fmt.Println(sections[i])
		log_format, _ := cfg.GetValue(sections[i], "log_format")
		rotate_type, _ := cfg.GetValue(sections[i], "rotate_type")
		//rotate_gap, _ := cfg.GetValue(sections[i], "rotate_gap")
		gsid_reg, _ := cfg.GetValue(sections[i], "gsid_reg")
		fmt.Println(gsid_reg)

		log_file := getLog(log_format)
		//for test
		//log_file = "access_log"

		config := tail.Config{
			ReOpen:	true,
			Follow:	true,
			Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
			MustExist: false,
			Poll:	  true,
			GsidReg:	gsid_reg,
		}

		logRotateConfig := tail.LogRotateConfig{
			LogFormat: log_format,
			LogRotateType: string(rotate_type),
			//LogRotateGap: int(rotate_gap),
			//LastRo
		}
		//fmt.Println(string(rotate_type))
		//fmt.Printf("%+v\n", logRotateConfig)

		//tailLog(ip, sections[i], log_file, config, logRotateConfig, app)
		go tailLog(ip, sections[i], log_file, config, logRotateConfig, app)
		//for test
		//break
	}

	<- log_chan

}
