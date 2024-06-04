package mylogger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

type LogLevel uint16

// MyLogger 定义一个接口
type MyLogger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

// 定义日志级别
const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// ParseLogLevel 将字符串类型转换为LogLevel类型
func ParseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":

		return TRACE, nil
	case "info":

		return INFO, nil
	case "warning":

		return WARNING, nil
	case "error":

		return ERROR, nil
	case "fatal":

		return FATAL, nil
	default:
		return UNKNOWN, nil
	}
}

// GetLogString 将LogLevel类型转换为字符串类型
func GetLogString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case TRACE:

		return "TRACE"
	case INFO:

		return "INFO"
	case WARNING:

		return "WARNING"
	case ERROR:

		return "ERROR"
	case FATAL:

		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// 找到语句调用的具体信息
func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed\n")
		return
	}
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	//按照"."切割funcName
	funcName = strings.Split(funcName, ".")[1]
	return
}
