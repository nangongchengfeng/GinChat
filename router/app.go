package router

import (
	"GinChat/docs"
	"GinChat/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 设置网页存放目录
	r.LoadHTMLGlob("views/**/*")
	// 设置静态文件存放目录
	r.Static("/asset", "./asset")
	r.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 首页
	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.GET("/toRegister", service.ToRegister)
	r.GET("/toChat", service.ToChat)
	r.POST("/searchFriends", service.SearchFriends)

	//用户模块
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("/user/createUser", service.CreateUser)
	r.GET("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)
	r.GET("/user/sendMsg", service.SendMsg)

	return r
}
