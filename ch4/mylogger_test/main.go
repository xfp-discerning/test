package main

import (
	"data/ch4/mylogger"
	"time"
)

//测试我们写的日志库的程序
func main() {
	// var str string
	// fmt.Scanln(str)
	for {
		log := mylogger.Newlog("debug")
		id := 100
		name := "xfp"
		log.Debug("这是一条Debug,id=%d,name=%s", id, name)
		log.Error("这是一条error")
		log.Trace("这是一条trace")
		log.Warning("这是一条warning")
		log.Info("这是一条info")
		log.Fatal("这是一条fatal")
		time.Sleep(time.Second)
	}
}
