package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 是一个全局数据库连接对象
var DB *gorm.DB

// Init 初始化数据库连接
func Init(dsn string) {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 在这里可以做自动迁移等操作
	// db.AutoMigrate(&models.User{})
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
