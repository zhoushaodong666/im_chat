package main

import (
	"im_chat/handle"
	"net/http"
)

func main() {
	//路由绑定
	//用户登录
	http.HandleFunc("/user/login", handle.UserLogin)

	//启动web服务器
	http.ListenAndServe(":8080", nil)
}
