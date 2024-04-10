package service

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles(
		"./views/chat/index.html", // 调整为正确的路径
		"./views/chat/head.html",  // 调整为正确的路径
	)
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "index")
	//c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})

}

func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "register")
	// c.JSON(200, gin.H{
	// "message": "welcome !! ",
	// })
}
