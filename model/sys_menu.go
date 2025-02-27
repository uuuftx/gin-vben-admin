package models

import (
	"time"
)

type SysMenu struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	ParentID  int       `gorm:"default:0" json:"parent_id"` // 父菜单 ID，0 表示顶级菜单
	Name      string    `gorm:"size:50;unique;not null" json:"name"`
	Path      string    `gorm:"size:100;not null" json:"path"`
	Component string    `gorm:"size:100" json:"component"` // 组件路径（可选）
	Icon      string    `gorm:"size:50" json:"icon"`       // 菜单图标（可选）
	Title     string    `gorm:"size:100;not null" json:"title"`
	Sort      int       `gorm:"default:0" json:"sort"` // 排序字段
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 设置 GORM 使用的数据库表名
func (SysMenu) TableName() string {
	return "sys_menu"
}

// 定义前端需要的菜单结构
type FrontendMenu struct {
	Name      string         `json:"name"`
	Path      string         `json:"path"`
	Component string         `json:"component,omitempty"`
	Meta      FrontendMeta   `json:"meta"`
	Children  []FrontendMenu `json:"children,omitempty"`
}

type FrontendMeta struct {
	Icon  string `json:"icon"`
	Title string `json:"title"`
}

// BuildMenuTree 将数据库菜单转换为前端需要的树形结构
func BuildMenuTree(menus []SysMenu, parentID int) []FrontendMenu {
	var tree []FrontendMenu

	for _, menu := range menus {
		if menu.ParentID == parentID {
			// 构建当前菜单项
			frontendMenu := FrontendMenu{
				Name:      menu.Name,
				Path:      menu.Path,
				Component: menu.Component,
				Meta: FrontendMeta{
					Icon:  menu.Icon,
					Title: menu.Title,
				},
				Children: BuildMenuTree(menus, menu.ID), // 递归构建子菜单
			}

			// 添加到树中
			tree = append(tree, frontendMenu)
		}
	}

	return tree
}
