package main

import (
	"imooc/crawer/engine"
	"imooc/crawer/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}