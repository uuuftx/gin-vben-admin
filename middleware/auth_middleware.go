package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/uuuftx/gin-vben-admin/utils"
	"strings"
)

// AuthMiddleware 验证 Authorization 请求头
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果是登录接口，直接跳过认证
		if c.Request.URL.Path == "/api/v1/auth/login" {
			c.Next() // 跳过中间件，继续执行
		}

		// 获取请求头中的 Authorization 字段
		authHeader := c.GetHeader("Authorization")

		// 如果 Authorization 字段为空，返回 401 错误
		if authHeader == "" {
			utils.Error(c, 401, "No Authorization")
			c.Abort() // 终止请求
		}

		// Authorization 字段格式应该为 "Bearer token"
		// 检查字段是否以 "Bearer " 开头
		if !strings.HasPrefix(authHeader, "Bearer ") {
			utils.Error(c, 401, "Authorization header must start with 'Bearer '")
			c.Abort() // 终止请求
		}

		// 获取 token 字符串，去掉 "Bearer " 前缀
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// 校验 token（可以替换为你自己的校验逻辑）
		if !isValidToken(token) {
			utils.Error(c, 401, "The provided token is invalid or expired")
			c.Abort() // 终止请求
		}

		// 如果 token 校验通过，继续执行后续的处理
		c.Next()
	}
}

// isValidToken 校验 token 是否有效（这个方法可以根据实际需求修改）
func isValidToken(token string) bool {
	// 在这里可以实现自己的 token 校验逻辑，举例来说可以是解析 JWT，查询数据库等。
	// 这里为了简化，假设有效 token 为 "valid-token"。
	user, err := utils.ParseToken(token)
	if err != nil {
		return false
	}
	if user == nil {
		return false
	}
	return true
}
