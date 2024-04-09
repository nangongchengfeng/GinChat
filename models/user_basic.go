package models

import (
	"GinChat/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// 定义一个UserBasic结构体，
type UserBasic struct {
	gorm.Model
	// 名字
	Name string
	// 密码
	PassWord string
	// 手机号
	Phone string
	// 邮箱
	Email string
	// 身份
	Identity string
	// 客户端IP
	ClientIp string
	// 客户端端口
	ClientPort string
	// 登录时间
	LoinTime time.Time
	// 心跳时间
	HeartbeatTime time.Time
	// 登录登出时间
	LoginOutTime time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	// 是否登出
	IsLogout bool
	// 设备信息
	DeviceInfo string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

// GetUserList 定义一个函数GetUserList，返回一个用户基本信息的切片
func GetUserList() []*UserBasic {
	// 创建一个用户基本信息的切片
	data := make([]*UserBasic, 10)
	// 从数据库中查找data中的数据
	utils.DB.Find(&data)
	// 遍历data中的数据
	for _, v := range data {
		// 打印v
		fmt.Println(v)
	}
	// 返回data
	return data
}
