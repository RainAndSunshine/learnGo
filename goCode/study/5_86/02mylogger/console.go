package mylogger

import (
	"fmt"
	"time"
)

//往终端写日志相关内容

// Logger 日志结构体
type Logger struct {
	Level LogLevel
}

// NewLog 构造函数
func NewLog(levelStr string) Logger {
	level, err := ParseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (l Logger) enable(logLevel LogLevel) bool {
	return l.Level <= logLevel
}

func (l Logger) log(level LogLevel, format string, a ...interface{}) {
	if l.enable(level) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), GetLogString(level), funcName, fileName, lineNo, msg)
	}
}

func (l Logger) Debug(format string, a ...interface{}) {
	l.log(DEBUG, format, a...)
}

func (l Logger) Info(format string, a ...interface{}) {
	l.log(INFO, format, a...)
}

func (l Logger) Warning(format string, a ...interface{}) {
	l.log(WARNING, format, a...)
}

func (l Logger) Error(format string, a ...interface{}) {
	l.log(ERROR, format, a...)
}

func (l Logger) Fatal(format string, a ...interface{}) {
	l.log(FATAL, format, a...)
}
