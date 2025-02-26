package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`    // 状态码
	Data    interface{} `json:"data"`    // 返回数据
	Message string      `json:"message"` // 返回消息
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code:    0,
		Data:    data,
		Message: "Success",
	})
}

// Error 错误响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Data:    nil,
		Message: message,
	})
}
