package main

import (
	"GinChat/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 测试代码
	db, err := gorm.Open(mysql.Open("root:123456@tcp(192.168.102.20:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 迁移 schema
	dbErr := db.AutoMigrate(&models.UserBasic{})
	if dbErr != nil {
		panic(dbErr)
	}
	// 创建记录
	user := &models.UserBasic{}
	user.Name = "深寒"
	db.Create(user)

	//查询记录
	fmt.Println(db.First(user, 1))

	// Update - 将 product 的 price 更新为 200
	db.Model(user).Update("PassWord", "1234")
}
