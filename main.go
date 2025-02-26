package main

import (
	"github.com/uuuftx/gin-vben-admin/router"
)

func main() {
	// 初始化路由
	r := router.SetupRouter()

	// 启动服务器
	r.Run(":8080")
}
