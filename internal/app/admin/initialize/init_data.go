package initialize

import (
	"lime/internal/app/admin/model"

	adapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func InitAdminData(db *gorm.DB) {
	// 初始化数据库
	InitAutoMigrate(db)

	// 初始化用户
	InitAdminUser()

	// 初始化菜单
	InitMenus()

	// 初始化接口
	InitApis()

	// 初始化角色
	InitRoles()
	InitUserRole()
	InitRoleMenus()
	InitRoleApi()

	// 初始化字典
	InitDictData()
}

func InitAutoMigrate(db *gorm.DB) {
	// 自动迁移模式
	db.AutoMigrate(model.User{})         // 用户信息
	db.AutoMigrate(model.Role{})         // 角色信息
	db.AutoMigrate(model.UserRole{})     // 用户角色信息
	db.AutoMigrate(model.RoleMenu{})     // 角色菜单信息
	db.AutoMigrate(adapter.CasbinRule{}) // 权限信息
	db.AutoMigrate(model.ApiInfo{})      // 接口信息
	db.AutoMigrate(model.Menu{})         // 菜单信息
	db.AutoMigrate(model.DictsInfo{})    // 字典信息
	db.AutoMigrate(model.DictDetail{})   // 字典详情信息
}
