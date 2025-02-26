package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/uuuftx/gin-vben-admin/utils"
)

func GetUsers(c *gin.Context) {
	// 模拟返回用户列表
	users := []string{"User1", "User2", "User3"}
	utils.Success(c, users)
}

func CreateUser(c *gin.Context) {
	// 模拟创建用户
	utils.Success(c, "User created")
}
