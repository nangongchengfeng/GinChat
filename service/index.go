package service

import (
	"GinChat/models"
	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

// GetUserList
// @Tags 用户列表信息
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	date := make([]*models.UserBasic, 10)
	date = models.GetUserList()

	c.JSON(
		200, gin.H{
			"data": date,
		})
}
