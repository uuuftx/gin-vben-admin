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
		apiV1.GET("/users", v1.GetUsers)
		apiV1.POST("/users", v1.CreateUser)
	}

	return r
}
