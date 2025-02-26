package v1

import (
	"github.com/gin-gonic/gin"
	models "github.com/uuuftx/gin-vben-admin/model"
	"github.com/uuuftx/gin-vben-admin/service"
	"github.com/uuuftx/gin-vben-admin/utils"
)

// GetUsers 返回用户列表
func GetUsers(c *gin.Context) {
	// 调用服务层获取用户列表
	users, err := service.GetUsers()
	if err != nil {
		// 如果查询失败，返回错误响应
		utils.Error(c, 100, err.Error())
		return
	}

	// 如果查询成功，返回用户列表
	utils.Success(c, users)
}

// CreateUser 接收用户数据并调用服务层创建用户
func CreateUser(c *gin.Context) {
	// 定义结构体来接收前端传递的用户数据
	var user models.SysUser

	// 解析前端请求的 JSON 数据到 user 结构体
	if err := c.ShouldBindJSON(&user); err != nil {
		// 如果解析失败，返回错误响应
		utils.Error(c, 100, err.Error())
		return
	}

	// 调用服务层创建用户的函数
	if err := service.CreateUser(&user); err != nil {
		// 如果创建失败，返回错误响应
		utils.Error(c, 101, err.Error())
		return
	}

	// 成功响应
	utils.Success(c, "User created successfully")
}
