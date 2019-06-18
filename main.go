package main

import (
	"io"
	"net/http"
)

func main() {
	//路由绑定
	http.HandleFunc("/user/login", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world!")
		//数据库操作
		//逻辑处理
		//rest api json 返回
		//1.获取前端传递的参数
		//moblie passwd
		//解析参数
		r.ParseForm()
		//如何获得参数
		mobile := r.PostForm.Get("mobile")
		passwd := r.PostForm.Get("passwd")

		loginok := false

		if mobile == "13715277993" && passwd == "123456" {
			loginok = true
		}
		str := `{"code":0,"data":{"id":1,"token":"test"}}`
		if !loginok {
			//成功

		}

			//设置header 为JSON 默认的 text/html 所以特别指定返回个格式
			//applicaltion/json
			w.Header().Set("Content-Type", "applicaltion/json")
			//设置状态码200
			w.WriteHeader(http.StatusOK)
			//写入ResponseWriter
			w.Write([]byte(str))
		else {
			//失败
			str := `{"code":-1,"msg":"密码不正确"}`
			//设置header 为JSON 默认的 text/html 所以特别指定返回个格式
			//applicaltion/json
			w.Header().Set("Content-Type", "applicaltion/json")
			//设置状态码200
			w.WriteHeader(http.StatusOK)
			//写入ResponseWriter
			w.Write([]byte(str))
		}
		//如何返回json

	})

	//启动web服务器
	http.ListenAndServe(":8080", nil)
}
