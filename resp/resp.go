package resp

import (
	"encoding/json"
	"log"
	"net/http"
)

type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	//设置状态码200
	w.WriteHeader(http.StatusOK)
	//写入ResponseWriter
	//定义一个结构体
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	//将结构体转化为json
	ret, err := json.Marshal(&h)
	if err != nil {
		log.Println(err.Error())
	}
	w.Write(ret)
}
