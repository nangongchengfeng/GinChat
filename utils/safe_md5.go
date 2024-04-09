package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

// 计算MD5哈希值
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// 计算MD5哈希值并转换为大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 生成密码
func MakePassword(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

// 解密
func ValidPassword(plainpwd, salt string, password string) bool {
	md := Md5Encode(plainpwd + salt)
	fmt.Println(md + " " + password)
	return md == password
}

//
//func main() {
//	date := "123456"
//	fmt.Println(MD5Encode(date)) // 输出：e10adc3949ba59abbe56e057f20f883e
//}
