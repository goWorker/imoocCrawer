package main

import (
	"imoocmp/crawer/engine"
	"imoocmp/crawer/scheduler"
	"imoocmp/crawer/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount:10,}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}