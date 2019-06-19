package main

import (
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"im_chat/controller"
	"log"
	"net/http"
)

//注册模板页面
func ResiterView() {
	// **表示目录 *表示文件
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		//打印错误并直接退出
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		http.HandleFunc(tplname, func(writer http.ResponseWriter, request *http.Request) {
			tpl.ExecuteTemplate(writer, tplname, nil)
		})
	}

}

func main() {
	//路由绑定
	//用户登录
	http.HandleFunc("/user/login", controller.UserLogin)
	//用户注册
	http.HandleFunc("/user/register", controller.UserRegister)

	http.HandleFunc("/contact/loadcommunity", controller.LoadCommunity)
	http.HandleFunc("/contact/loadfriend", controller.LoadFriend)
	http.HandleFunc("/contact/joincommunity", controller.JoinCommunity)

	http.HandleFunc("/contact/addfriend", controller.Addfriend)
	http.HandleFunc("/chat", controller.Chat)

	//1.提供静态资源目录支持
	//http.Handle("/",http.FileServer(http.Dir(".")))
	//2.提供指定目录的静态文件支持
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	ResiterView()
	//启动web服务器
	http.ListenAndServe(":8080", nil)
}
