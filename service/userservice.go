package service

import (
	"GinChat/models"
	"GinChat/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"strconv"
)

const salt = "heian"

// GetUserList
// @Summary 所有用户
// @Tags  用户模块
// @Success 200 {string} json {"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	date := make([]*models.UserBasic, 10)
	date = models.GetUserList()

	c.JSON(
		200, gin.H{
			"message": date,
		})
}

// FindUserByNameAndPwd
// @Summary 所有用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {

	//data := models.UserBasic{}
	name := c.Query("name")
	password := c.Query("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"message": "用户名不存在！",
		})
		return
	}
	flag := utils.ValidPassword(password, salt, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"message": "密码不正确",
		})
		return
	}
	pwd := utils.MakePassword(password, salt)
	data := models.FindUserByNameAndPwd(name, pwd)
	c.JSON(200, gin.H{
		"code":    0, // 0成功 -1失败
		"message": "查询成功！",
		"data":    data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags  用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json {"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	data := models.FindUserByName(user.Name)

	if data.Name != "" {
		c.JSON(-1, gin.H{
			"message": "用户名已注册！",
		})
		return
	}

	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "密码不一致",
		})
		return
	}
	// 进行密码加密
	user.PassWord = utils.MakePassword(password, salt)
	//user.PassWord = password
	models.CreateUser(user)
	c.JSON(
		200, gin.H{
			"message": "新增用户成功！",
		})
}

// DeleteUser
// @Summary 删除用户
// @Tags  用户模块
// @param id query string false "用户id"
// @Success 200 {string} json {"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除用户成功！"})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = utils.MakePassword(c.PostForm("password"), salt)
	//user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	fmt.Println("update:", user)
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(200, gin.H{
			"message": "修改参数不匹配！",
		})
	} else {
		models.UpdateUser(user)
		//c.JSON(200, gin.H{
		//	"message": "修改用户成功！",
		//})
		c.JSON(200, gin.H{
			"code":    0, // 0成功 -1失败
			"message": "修改用户成功！",
			"data":    user,
		})
	}

}