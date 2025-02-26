package main

import (
	"fmt"
	"github.com/uuuftx/gin-vben-admin/config"
	"github.com/uuuftx/gin-vben-admin/db"
	"github.com/uuuftx/gin-vben-admin/router"
)

func main() {
	// 加载配置文件
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	// 打印配置
	fmt.Printf("Server Port: %d\n", cfg.Server.Port)
	fmt.Printf("Database Host: %s\n", cfg.Database.Host)

	// 数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)
	db.Init(dsn) // 初始化数据库连接

	// 初始化路由
	r := router.SetupRouter()

	// 启动服务器
	err = r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		return
	}
}
