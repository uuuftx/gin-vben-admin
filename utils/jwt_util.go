package utils

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

// 定义 JWT 密钥，实际项目中应该放在环境变量或者配置文件中
var jwtSecretKey = []byte("your-secret-key")

// JWTClaims 自定义的 JWT 声明
type JWTClaims struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

// GenerateToken 生成 JWT
func GenerateToken(userID, userName string) (string, error) {
	// 设置过期时间（例如1小时）
	expirationTime := time.Now().Add(1 * time.Hour)

	// 创建 JWT 声明
	claims := &JWTClaims{
		UserID:   userID,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "gin-vben-admin", // 你可以自定义 Issuer
		},
	}

	// 创建一个 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名生成 token 字符串
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析 JWT
func ParseToken(tokenStr string) (*JWTClaims, error) {
	// 解析 token
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 这里只需要返回签名时使用的密钥
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// 如果 token 有效且解析成功，返回 claims
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// GetUserInfo 从请求的 Authorization header 中提取并解析 JWT，返回用户信息
func GetUserInfo(authHeader string) (*JWTClaims, error) {

	// 解析 token，获取用户信息
	token := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	// 返回解析后的用户信息
	return claims, nil
}
