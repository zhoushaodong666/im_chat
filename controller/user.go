package controller

import (
	"fmt"
	"im_chat/model"
	"im_chat/service"
	"im_chat/util"
	"math/rand"
	"net/http"
)

var userService service.UserService

//用户登录处理函数
func UserLogin(w http.ResponseWriter, r *http.Request) {
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

	user, err := userService.Login(mobile, passwd)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, user, "")
	}
}

//用户注册处理函数
func UserRegister(w http.ResponseWriter, r *http.Request) {
	//1.解析参数
	r.ParseForm()
	//2.获取参数
	mobile := r.PostForm.Get("mobile")
	plainpwd := r.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKONW
	//3.执行注册逻辑
	user, err := userService.Register(mobile, plainpwd, nickname, avatar, sex)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, user, "")
	}
}
