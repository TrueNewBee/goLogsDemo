package main

import (
	"time"

	"chentianxiang.vip/studygo/day06/mylogger"
)

// 测试我们自己 写的日志库
// 调用了mylogger包
var log mylogger.Logger //声明一个全局的接口变量

// 选择日志方式
func choice(i int) {
	switch i {
	case 1: // 终端输出
		log = mylogger.NewConsoLoger("Info")
	case 2: // 文件记录
		log = mylogger.NewFileLogger("Info", "./", "zhoulinwan.log", 10*1024*1024) // 文件日志输出示例
	default:
		return
	}
}
func main() {
	// 选择哪种日志方式
	choice(2)
	// log = mylogger.NewConsoLoger("Info")                                       // 终端日志输出示例
	// log = mylogger.NewFileLogger("Info", "./", "zhoulinwan.log", 10*1024*1024) // 文件日志输出示例
	for {
		id := 10010
		name := "理想"
		log.Debug("这是一条Debug日志, id:%d, name:%s", id, name)
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志, id:%d, name:%s", id, name)
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second * 1)
	}
}
