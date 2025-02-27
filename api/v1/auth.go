package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/uuuftx/gin-vben-admin/service"
	"github.com/uuuftx/gin-vben-admin/utils"
)

func Login(c *gin.Context) {
	// 获取请求中的用户名和密码（假设请求体是 JSON 格式）
	var loginData struct {
		Captcha       bool   `json:"captcha"`
		Username      string `json:"username"`
		Password      string `json:"password"`
		SelectAccount string `json:"selectAccount"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		utils.Error(c, 100, err.Error())

	}
	user, err := service.FindUserByUsername(loginData.Username)

	if err != nil {
		utils.Error(c, 200, "无此用户")
		return
	}

	if user == nil {
		utils.Error(c, 200, "没有此用户")
		return
	}

	passwordReq := utils.MD5Encrypt(loginData.Password)
	if passwordReq != user.Password {
		utils.Error(c, 100, "密码错误")
		return
	}

	// 用户名和密码验证通过，生成 JWT
	token, err := utils.GenerateToken(user.UserID, user.UserName)
	if err != nil {
		utils.Error(c, 100, "jwt生成失败")
		return
	}

	// 这里可以实现你的用户名密码验证逻辑
	// 如果验证成功，返回登录成功的消息
	utils.Success(c, gin.H{
		"username":    user.UserName,
		"accessToken": token,
		"realName":    user.RealName,
	})
}
