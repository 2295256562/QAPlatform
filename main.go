package main

import (
	"QAPlatform/model"
	"QAPlatform/routers"
)

func main() {
	// 加载数据库
	model.InitDb()
	routers.InitRouter()

}