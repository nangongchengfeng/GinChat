package service

import (
	"GinChat/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles(
		"./views/user/index.html", // 调整为正确的路径
		"./views/chat/head.html",  // 调整为正确的路径
	)
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "index")
	//c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})

}

func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/register.html",
		"./views/chat/head.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "register")
	// c.JSON(200, gin.H{
	// "message": "welcome !! ",
	// })
}

func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles("views/chat/index.html",
		"views/chat/head.html",
		"views/chat/foot.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/main.html")
	if err != nil {
		panic(err)
	}
	userId, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(userId)
	user.Identity = token
	//fmt.Println("ToChat>>>>>>>>", user)
	ind.Execute(c.Writer, user)
}
