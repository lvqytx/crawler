package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	// 串行
	// engine.Run(engine.Request{	// 配置请求信息即可
	// 	Url:       "http://www.zhenai.com/zhenghun",
	// 	ParseFunc: parser.ParseCityList,
	// })

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 50,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
