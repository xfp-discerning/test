package mylogger

import (
	"fmt"
	"time"
)

//往终端写相关日志文件
type Logger struct {
	Level LogLevel
}

func Newlog(Levelstr string) Logger {
	level, err := parseLevelstr(Levelstr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (l Logger) enable(loglevel LogLevel) bool {
	return loglevel >= l.Level
}

func log(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	funcName, fileName, lineNo := getInfo(3)
	now := time.Now().Format("2006/01/02 15:04:05")
	fmt.Printf("[%s]   [%s]    [%s::%s::%d]\n", now, msg, funcName, fileName, lineNo)
}

func (l Logger) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		log(format, a...)
	}

}
func (l Logger) Trace(format string, a ...interface{}) {
	if l.enable(TRACE) {
		log(format, a...)
	}
}

func (l Logger) Info(format string, a ...interface{}) {
	if l.enable(INFO) {
		log(format, a...)
	}
}
func (l Logger) Warning(format string, a ...interface{}) {
	if l.enable(WARNING) {
		log(format, a...)
	}
}

func (l Logger) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		log(format, a...)
	}
}

func (l Logger) Fatal(format string, a ...interface{}) {
	if l.enable(FATAL) {
		log(format, a...)
	}
}
