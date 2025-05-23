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

	// 区分 GET 和 POST 请求
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintf(w, "GET GoLift!")
		}
		if r.Method == http.MethodPost {
			fmt.Fprintf(w, "POST GoLift!")
		}
	})

	// 启动服务
	http.ListenAndServe(":"+strconv.Itoa(base.Conf.Http.Port), nil)
}
