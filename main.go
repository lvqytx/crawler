package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{	// 配置请求信息即可
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
	
}
