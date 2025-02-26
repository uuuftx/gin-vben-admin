package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
	"strings"
	"time"
)

// GenerateUUIDWithoutDash 生成一个没有短横线的 UUID
func GenerateUUIDWithoutDash() string {
	// 生成 UUID
	uuidWithDash := uuid.New()

	// 转换为字符串并去掉短横线
	return strings.ReplaceAll(uuidWithDash.String(), "-", "")
}

// MD5Encrypt 对输入字符串进行 MD5 加密
func MD5Encrypt(input string) string {
	// 计算 MD5 哈希值
	hash := md5.New()
	hash.Write([]byte(input))

	// 获取哈希结果并转换为十六进制字符串
	hashInBytes := hash.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}

// GetCurrentTime 返回当前时间
func GetCurrentTime() string {
	// 获取当前时间并格式化为字符串（年月日时分秒）
	return time.Now().Format("2006-01-02 15:04:05")
}
