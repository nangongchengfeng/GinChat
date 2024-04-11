package models

import (
	"GinChat/utils"
	"fmt"
	"gorm.io/gorm"
)

// 人员关系
type Contact struct {
	gorm.Model
	OwnerId  uint //谁的关系信息
	TargetId uint //对应的谁
	Type     int  //对应的类型 0 1 3
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}

// 搜索好友
func SearchFriend(userId uint) []UserBasic {
	// 创建联系人切片
	contacts := make([]Contact, 0)
	// 创建对象ID切片
	objIds := make([]uint64, 0)
	// 在数据库中查找所有属于userId的用户
	utils.DB.Where("owner_id = ? and type = 1", userId).Find(&contacts)
	// 遍历联系人切片
	for _, v := range contacts {
		fmt.Println(">>>>", v)
		// 将目标ID添加到对象ID切片中
		objIds = append(objIds, uint64(v.TargetId))
	}
	// 创建用户切片
	users := make([]UserBasic, 0)
	// 在数据库中查找所有属于对象ID切片中的用户
	utils.DB.Where("id in (?)", objIds).Find(&users)
	// 返回用户切片
	return users
}
