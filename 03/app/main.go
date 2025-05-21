package main

import (
	"base"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// 初始化配置
	base.InitConfig("./conf.yml")

	// 第一个HTTP接口
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello GoLift!")
	})

	// 启动服务
	http.ListenAndServe(":"+strconv.Itoa(base.Conf.Http.Port), nil)
}
