package handle

import (
	"im_chat/resp"
	"io"
	"net/http"
)

//用户登录处理函数
func UserLogin(w http.ResponseWriter, r *http.Request) {
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
	if loginok {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		resp.Resp(w, 0, data, "")
	} else {
		resp.Resp(w, -1, nil, "密码不正确")
	}
	//设置header 为JSON 默认的 text/html 所以特别指定返回个格式
	//application/json

}
