package v1

import (
	"github.com/uuuftx/gin-vben-admin/db"
	models "github.com/uuuftx/gin-vben-admin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMenus 获取所有菜单
func GetMenus(c *gin.Context) {
	var menus []models.SysMenu
	if err := db.DB.Find(&menus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menus)
}

// GetMenuTree 获取菜单树
func GetMenuTree(c *gin.Context) {
	var menus []models.SysMenu
	if err := db.DB.Order("sort ASC").Find(&menus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 转换为树形结构
	menuTree := models.BuildMenuTree(menus, 0)

	c.JSON(http.StatusOK, menuTree)
}

// CreateMenu 创建菜单
func CreateMenu(c *gin.Context) {
	var menu models.SysMenu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menu)
}

// UpdateMenu 更新菜单
func UpdateMenu(c *gin.Context) {
	id := c.Param("id")
	var menu models.SysMenu
	if err := db.DB.First(&menu, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menu)
}

// DeleteMenu 删除菜单
func DeleteMenu(c *gin.Context) {
	id := c.Param("id")
	var menu models.SysMenu
	if err := db.DB.First(&menu, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}

	if err := db.DB.Delete(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu deleted"})
}
