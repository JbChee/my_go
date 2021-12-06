package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"log"
	"os"
)



func test1(){

	fmt.Println(".........")	
	log.Println("test1  dfsaf")

	v := "string"
	log.Printf("log  %s  a a \n",v)

	log.Fatalln("this is a fatal log ")
	log.Panicln("this is a panic log")
}

//设置日志
func setlog(){

	
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println(" 这是一条普通日志")


	log.SetPrefix("[***]")
	log.Println(" 这是一条普通日志1111111")

	
	///[***]2020/08/14 00:02:29.494603 E:/Go_code/src/log_deman/logger_deman.go:13: test1  dfsaf
	///[***]2020/08/14 00:02:29.496602 E:/Go_code/src/log_deman/logger_deman.go:16: log  string  a a
	///[***]2020/08/14 00:02:29.497605 E:/Go_code/src/log_deman/logger_deman.go:18: this is a fatal log  
	


}

//写入日志到文件中
func writelogfile1(filepath string) bool{

	
	w := false

	logFile, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil  {
		fmt.Println("open log file failed , err:",err)
		return w
	}else{
		log.SetOutput(logFile)
		setlog()
		w = true
		return w
	
	}

}

//log 读取配置文件
func read_log_conf(logfilepath string){

	logconf, err := config.NewConfig("ini", logfilepath)
	if err != nil {
		fmt.Println("new config failed , err:", err)
		return 
	}
	fmt.Println(logconf)
	port,err := logconf.Int("server::listen_port")
	if err != nil {
		fmt.Println("read server:port failed,err:", err)
		return 
	}
	fmt.Println("port:",port)

	log_level := logconf.String("logs::log_level")
	if len(log_level) == 0 {
		log_level = "debug"
	}

	fmt.Println("log_level:",log_level)

	log_path := logconf.String("collect::log_path")
	fmt.Println("log_path:",log_path)


}

//本地配置 config = {}
func read_log_conf2(logfilepath string){

	config :=make(map[string]interface{})

	config["filename"] = logfilepath
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
        fmt.Println("marshal failed",err)
        return
	}

	fmt.Printf("configstr:%v",configStr)
	//configstr  json数据
	logs.SetLogger(logs.AdapterFile, string(configStr))

    logs.Debug("this is a test,my name is %s","stu01")
    logs.Trace("this is a trace")
    logs.Warn("this is a warn")

}

func main(){

	// setlog()

	logpath1 :="../logs/test_log.txt"
	w :=writelogfile1(logpath1)

	fmt.Println(w)
	// test1()

	
	//log conf path
	conf_path := "../logs/logconf.conf"
	read_log_conf(conf_path)

	//log 本地配置
	logpath2 :="E:\\Go_code\\src\\logs\\logcollect.log"
	read_log_conf2(logpath2)

}
