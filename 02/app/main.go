package main

import (
	"base"
	"fmt"
)

func main() {
	base.InitConfig("./conf.yml")
	fmt.Printf("Http端口：%d\n", base.Conf.Http.Port)
	fmt.Printf("数据库路径：%s\n", base.Conf.Db.Path)
	fmt.Println("Hello GoLift!")
}
