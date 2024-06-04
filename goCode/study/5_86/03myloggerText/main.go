package main

import (
	mylogger "goCode/study/5_86/02mylogger"
)

//自定义一个日志库

//需求分析
//1.支持往不同地方输出日志
//2.日志分级别
//		2.1 Debug
//		2.2 Trace
//		2.3 Info
//  	2.4 Warning
//  	2.5 Error
//  	2.6 Fatal
//  	...
//3.日志要支持开关控制,比如说开发的时候所有级别都能输出，但是上线之后只有INFO级别往下的才能输出
//4.一个完整的日志记录要包含：时间，行号，文件名，日志级别，日志信息
//5.日志文件要切割
//		5.1 按文件大小切割
//			1)每次记录日志之前都判断一下当前写的这个文件的文件大小
//		5.2 按日期切割
//			1)在日志结构体中设置一个字段记录上一次切割的小时数
//			2)在写日志之前检查当前时间的小时数和之前保存的是否一致，不一致就需要切割

//测试我们自己写的日志库的程序

// 声明一个全局的接口变量
var log mylogger.MyLogger

func main() {
	log = mylogger.NewLog("debug")                                           //终端日志
	log = mylogger.NewFileLogger("debug", "./", "zhoulin.log", 10*1024*1024) //文件日志
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		id := 10010
		name := "理想"
		log.Error("这是一条Error日志,id:%d ,name:%s", id, name)
		log.Fatal("这是一条Fatal日志")
		//time.Sleep(2 * time.Second)
	}
}
