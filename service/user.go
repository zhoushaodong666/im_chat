package service

import (
	"fmt"
	"github.com/pkg/errors"
	"im_chat/model"
	"im_chat/util"
	"math/rand"
	"time"
)

type UserService struct {
}

//用户注册函数
func (s *UserService) Register(mobile, plainpwd, nickname, avatar, sex string) (user model.User, err error) {
	//检测手机号是否存在
	tmp := model.User{}
	_, err = DbEngine.Where("mobile=?", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	//如果存在则返回提示已经注册
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已被注册")
	}
	//否则插入数据
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	//加密密码
	tmp.Passwd = util.MakePasswd(plainpwd, tmp.Salt)
	tmp.Createat = time.Now()
	//token  可以是个随机数
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())
	//插入数据库
	_, err = DbEngine.InsertOne(&tmp)

	//返回新用户信息
	return tmp, nil
}

//用户登陆函数
func (s *UserService) Login(mobile, plainpwd string) (user model.User, err error) {
	tmp := model.User{}
	_, err = DbEngine.Where("mobile = ?", mobile).Get(&tmp)
	//用户不存在
	if tmp.Id == 0 {
		return tmp, errors.New("该用户不存在")
	}
	//比对密码
	if !util.ValidatePasswd(plainpwd, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("密码错误")
	}
	//刷新token 安全性
	str := fmt.Sprintf("%d", time.Now().UnixNano())
	token := util.MD5Encode(str)
	tmp.Token = token
	DbEngine.ID(tmp.Id).Cols("token").Update(&tmp)

	return tmp, nil
}

//查找某个用户
func (s *UserService) Find(
	userId int64) (user model.User) {

	//首先通过手机号查询用户
	tmp := model.User{}
	DbEngine.ID(userId).Get(&tmp)
	return tmp
}
