package service

import (
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	commonModel "lime/internal/common/model"
	"lime/internal/common/service"
	"lime/internal/global"

	"gorm.io/gorm"
)

type User struct {
	service.BaseService[model.User]
}

func NewUser(args ...any) *User {
	srv := &User{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// 根据用户名查询用户
func (u *User) GetUserByName(name string) (*model.User, error) {
	user := &model.User{}
	err := global.DB.Model(&model.User{}).Where("username = ?", name).Preload("UserRole").Find(user).Error
	return user, err
}

// 根据条件查询列表
func (u *User) ListUser(condition *commonModel.PageQuery[*requests.QueryUser]) (*commonModel.ResPage[model.User], error) {
	return service.GetList[model.User](condition, func(qu *requests.QueryUser, d *gorm.DB) *gorm.DB {
		d = d.Preload("UserRole")

		if qu.KeyWords != "" {
			d = d.Where("username like ?", "%"+qu.KeyWords+"%")
		}

		return d
	})
}

// 添加用户
func (u *User) AddUser(user *model.User) error {
	return global.DB.Model(&model.User{}).Create(user).Error
}

// 更新用户
func (u *User) UpdateUser(user *model.User) error {
	return global.DB.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error
}

// 更新用户
func (u *User) UpdateUserByData(id uint, data map[string]any) error {
	return global.DB.Model(&model.User{}).Where("id = ?", id).Updates(data).Error
}

// 删除用户
func (u *User) DeleteUser(id int) error {
	return global.DB.Model(&model.User{}).Where("id = ?", id).Delete(&model.User{}).Error
}

// 根据用户id查询用户
func (u *User) GetUserById(id int) (*model.User, error) {
	user := &model.User{}
	err := global.DB.Model(&model.User{}).Where("id = ?", id).First(user).Error
	return user, err
}

// 根据用户id查询用户角色
func (u *User) GetUserRoleById(id int) (*model.Role, error) {
	role := new(model.Role)
	err := global.DB.Model(&model.Role{}).Joins("JOIN user_role ON user_role.Role_code = role.code").Where("user_role.user_id = ?", id).Find(role).Error
	return role, err
}

func (u *User) GetRoleByName(name string) (*model.Role, error) {
	role := new(model.Role)
	err := global.DB.Model(&model.Role{}).Where("name = ?", name).Find(role).Error
	return role, err
}

// AddRole
func (u *User) AddRole(role *model.Role) error {
	return global.DB.Model(&model.Role{}).Create(role).Error
}

// 查询角色菜单

func (u *User) CheckRoleMenu(code string, menu_code string) (bool, error) {
	count := int64(0)
	err := global.DB.Model(&model.RoleMenu{}).Where("role_code = ?", code).Where("menu_code = ?", menu_code).Count(&count).Error
	return count > 0, err
}

func (u *User) CheckUserRole(user_id int, role_code string) (bool, error) {
	count := int64(0)
	err := global.DB.Model(&model.UserRole{}).Where("user_id = ?", user_id).Where("role_code = ?", role_code).Count(&count).Error
	return count > 0, err
}

// AddUserRole
func (u *User) AddUserRole(userRole *model.UserRole) error {
	return global.DB.Model(&model.UserRole{}).Create(userRole).Error
}

// AddRoleMenu
func (u *User) AddRoleMenu(roleMenu *model.RoleMenu) error {
	return global.DB.Model(&model.RoleMenu{}).Create(roleMenu).Error
}
