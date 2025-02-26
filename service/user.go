package service

import (
	"fmt"
	"github.com/uuuftx/gin-vben-admin/db"
	models "github.com/uuuftx/gin-vben-admin/model"
	"github.com/uuuftx/gin-vben-admin/utils"
	"time"
)

func CreateUser(user *models.SysUser) error {
	user.UserID = utils.GenerateUUIDWithoutDash()
	user.Password = utils.MD5Encrypt(user.Password)
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	if err := db.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUsers() ([]models.SysUser, error) {
	var users []models.SysUser
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUsersWithConditionsAndPagination 根据条件和分页查询用户信息
func GetUsersWithConditionsAndPagination(page, pageSize int, userName, role string) ([]models.SysUser, int64, error) {
	var users []models.SysUser
	var query = db.DB.Model(&models.SysUser{})

	// 添加查询条件
	if userName != "" {
		query = query.Where("user_name LIKE ?", fmt.Sprintf("%%%s%%", userName))
	}
	if role != "" {
		query = query.Where("role = ?", role)
	}

	// 计算跳过的记录数
	offset := (page - 1) * pageSize

	// 执行分页查询
	if err := query.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	// 获取总记录数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 返回用户数据和总记录数
	return users, total, nil
}

func GetUserByID(id uint) (*models.SysUser, error) {
	var user models.SysUser
	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user *models.SysUser) error {
	if err := db.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uint) error {
	if err := db.DB.Delete(&models.SysUser{}, id).Error; err != nil {
		return err
	}
	return nil
}
