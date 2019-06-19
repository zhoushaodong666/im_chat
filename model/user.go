package model

import "time"

const (
	SEX_WOMEN  = "W"
	SEX_MAN    = "M"
	SEX_UNKONW = "U"
)

type User struct {
	//用户的ID
	Id int64 `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	//手机号码
	Mobile string `xorm:"varchar(20)" form:"mobile" json:"mobile"`
	//用户密码 = f(pwd+salt) MD5
	Passwd string `xorm:"varchar(40)" form:"passwd" json:"passwd"`
	//头像
	Avatar string `xorm:"varchar(150)" form:"avatar" json:"avatar"`
	//性别
	Sex string `xorm:"varchar(20)" form:"sex" json:"sex"`
	//昵称
	Nickname string `xorm:"varchar(20)" form:"nickname" json:"nickname"`
	//盐值 随机数
	Salt string `xorm:"varchar(10)" form:"salt" json:"salt"`
	//是否
	Online int `xorm:"int(10)" form:"online" json:"online"`
	//鉴权字段
	Token string `xorm:"varchar(40)" form:"token" json:"token"`

	Memo string `xorm:"varchar(140)" form:"memo" json:"memo"`
	//统计每天用户增量
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`
}
