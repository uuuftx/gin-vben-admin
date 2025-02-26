package models

import (
	"encoding/json"
	"time"
)

// SysUser 定义数据库中的 sys_user 表对应的结构体
type SysUser struct {
	ID          uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UserID      string    `gorm:"size:32;uniqueIndex;column:user_id" json:"user_id"` // 适当调整索引或唯一约束
	UserName    string    `gorm:"size:32;column:user_name" json:"user_name"`
	Password    string    `gorm:"size:64;column:password" json:"-"`
	RealName    string    `gorm:"size:32;column:real_name" json:"real_name"`
	Roles       string    `gorm:"size:255;column:roles" json:"roles"`
	Avatar      string    `gorm:"size:255;column:avatar" json:"avatar"`
	Description string    `gorm:"size:255;column:desc" json:"description"`
	HomePath    string    `gorm:"size:255;column:home_path" json:"home_path"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
}

// MarshalJSON 自定义 JSON 序列化
func (u SysUser) MarshalJSON() ([]byte, error) {
	// 创建一个匿名结构体，排除 Password 字段
	type Alias SysUser
	return json.Marshal(&struct {
		Password string `json:"-"`
		Alias
	}{
		Alias:    (Alias)(u),
		Password: "", // 隐藏密码字段
	})
}

// TableName 设置 GORM 使用的数据库表名
func (SysUser) TableName() string {
	return "sys_user"
}
