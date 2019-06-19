package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//返回小写MD5密文
func Md5Encode(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	encodeStr := hash.Sum(nil)
	return hex.EncodeToString(encodeStr)
}

//返回大写的MD5密文
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

//密文与数据库密码字段比较
func ValidatePasswd(plainpwd, salt, passwd string) bool {
	return Md5Encode(plainpwd+salt) == passwd
}

//返回一个 明文+盐值的md5字符串
func MakePasswd(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}
