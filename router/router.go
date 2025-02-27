package router

import (
	"github.com/gin-gonic/gin"
	"github.com/uuuftx/gin-vben-admin/api/v1"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 注册路由
	apiV1 := r.Group("/api/v1")
	{
		// 用户模块
		apiV1.GET("/user/info", v1.GetUser)
		apiV1.GET("/auth/codes", v1.AuthCode)

		//登录模块
		apiV1.POST("/auth/login", v1.Login)

		// 菜单模块
		apiV1.GET("/menu/all", v1.GetMenuTree)
	}

	return r
}
