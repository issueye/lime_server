package initialize

import (
	"lime/internal/app/admin/logic"
	"lime/internal/app/admin/model"
	"log/slog"

	"gorm.io/gorm"
)

func InitAdminData(db *gorm.DB) {
	InitAutoMigrate(db)

	InitRoles()
	InitMenus()
	InitApis()
}

func InitAutoMigrate(db *gorm.DB) {
	// 自动迁移模式
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Role{})
	db.AutoMigrate(&model.UserRole{})
	db.AutoMigrate(&model.RoleMenu{})
	db.AutoMigrate(&model.ApiInfo{})
	db.AutoMigrate(&model.Menu{})
	db.AutoMigrate(&model.DictsInfo{})
	db.AutoMigrate(&model.DictDetail{})

}

// 初始化角色数据
func InitRoles() {
	Roles := []model.RoleBase{
		{Code: "9001", Name: "管理员", IsCanNotRemove: 1, Remark: "系统生成"},
	}

	lc := logic.NewRoleLogic()
	for _, Role := range Roles {
		data := model.NewRole(Role)
		slog.Info("检查角色数据", "code", Role.Code, "name", Role.Name)
		lc.RoleIsNotExistAdd(data)
	}
}

// 初始化菜单数据
func InitMenus() {
	menus := []model.MenuBase{
		{Code: "1000", Name: "首页", Description: "首页", Frontpath: "/home", Order: 10, Visible: true, Icon: "Home", ParentCode: "", MenuType: model.EMT_MENU, IsLink: 0, IsCanNotRemove: 1},
		{Code: "9000", Name: "系统管理", Description: "系统管理", Frontpath: "/system", Order: 90, Visible: true, Icon: "Setting", ParentCode: "", MenuType: model.EMT_DIRECTORY, IsLink: 0, IsCanNotRemove: 1},
		{Code: "9001", Name: "用户管理", Description: "用户管理", Frontpath: "/system/user", Order: 91, Visible: true, Icon: "User", ParentCode: "9000", MenuType: model.EMT_MENU, IsLink: 0, IsCanNotRemove: 1},
		{Code: "9002", Name: "角色管理", Description: "角色管理", Frontpath: "/system/role", Order: 92, Visible: true, Icon: "Avatar", ParentCode: "9000", MenuType: model.EMT_MENU, IsLink: 0, IsCanNotRemove: 1},
		{Code: "9003", Name: "菜单管理", Description: "菜单管理", Frontpath: "/system/menu", Order: 93, Visible: true, Icon: "Menu", ParentCode: "9000", MenuType: model.EMT_MENU, IsLink: 0, IsCanNotRemove: 1},
		{Code: "9004", Name: "字典管理", Description: "字典管理", Frontpath: "/system/dict_mana", Order: 94, Visible: true, Icon: "List", ParentCode: "9000", MenuType: model.EMT_MENU, IsLink: 0, IsCanNotRemove: 1},
		{Code: "9005", Name: "系统设置", Description: "系统设置", Frontpath: "/system/setting", Order: 95, Visible: true, Icon: "Tools", ParentCode: "9000", MenuType: model.EMT_MENU, IsLink: 0, IsCanNotRemove: 1},
	}

	lc := logic.NewMenuLogic()
	for _, menu := range menus {
		m := model.BaseNewMenu(menu)
		slog.Info("检查菜单数据", "code", menu.Code, "name", menu.Name)
		lc.MenuIsNotExistAdd(m)
	}
}

func InitApis() {
	apis := []model.ApiBase{
		// 基础接口
		{Title: "登录接口", Path: "/api/v1/auth/login", Method: model.EnumApiMethodPost, Group: "基础接口"},
		{Title: "退出登录接口", Path: "/api/v1/auth/logout", Method: model.EnumApiMethodPost, Group: "基础接口"},
		{Title: "刷新token接口", Path: "/api/v1/auth/refresh", Method: model.EnumApiMethodGet, Group: "基础接口"},

		// 基础接口
		{Title: "获取用户信息接口", Path: "/api/v1/admin/userInfo", Method: model.EnumApiMethodGet, Group: "基础接口"},
		{Title: "上传文件接口", Path: "/api/v1/admin/upload", Method: model.EnumApiMethodPost, Group: "基础接口"},
		{Title: "更新用户信息接口", Path: "/api/v1/admin/updateuserinfo", Method: model.EnumApiMethodPost, Group: "基础接口"},
		{Title: "更新密码接口", Path: "/api/v1/admin/updatepassword", Method: model.EnumApiMethodPost, Group: "基础接口"},

		// 用户管理
		{Title: "根据条件查询数据", Path: "/api/v1/user/list", Method: model.EnumApiMethodPost, Group: "用户管理"},
		{Title: "修改用户信息", Path: "/api/v1/user/update", Method: model.EnumApiMethodPut, Group: "用户管理"},
		{Title: "删除用户信息", Path: "/api/v1/user/delete/{id}", Method: model.EnumApiMethodDelete, Group: "用户管理"},
		{Title: "创建用户信息", Path: "/api/v1/user/add", Method: model.EnumApiMethodPost, Group: "用户管理"},

		// 角色管理
		{Title: "根据条件查询数据", Path: "/api/v1/role/list", Method: model.EnumApiMethodPost, Group: "角色管理"},
		{Title: "修改角色信息", Path: "/api/v1/role/update", Method: model.EnumApiMethodPut, Group: "角色管理"},
		{Title: "删除角色信息", Path: "/api/v1/role/delete/{id}", Method: model.EnumApiMethodDelete, Group: "角色管理"},
		{Title: "创建角色信息", Path: "/api/v1/role/add", Method: model.EnumApiMethodPost, Group: "角色管理"},

		// 菜单管理
		{Title: "根据条件查询数据", Path: "/api/v1/menu/list", Method: model.EnumApiMethodPost, Group: "菜单管理"},
		{Title: "获取所有菜单", Path: "/api/v1/menu/getAll", Method: model.EnumApiMethodGet, Group: "菜单管理"},
		{Title: "获取角色菜单", Path: "/api/v1/menu/roleMenus/{code}", Method: model.EnumApiMethodGet, Group: "菜单管理"},
		{Title: "保存角色菜单", Path: "/api/v1/menu/saveRoleMenus/{code}", Method: model.EnumApiMethodPost, Group: "菜单管理"},
		{Title: "修改菜单", Path: "/api/v1/menu/update", Method: model.EnumApiMethodPut, Group: "菜单管理"},
		{Title: "删除菜单", Path: "/api/v1/menu/delete/{code}", Method: model.EnumApiMethodPut, Group: "菜单管理"},
		{Title: "创建菜单", Path: "/api/v1/menu/add", Method: model.EnumApiMethodPut, Group: "菜单管理"},

		// 系统设置
		{Title: "获取系统设置", Path: "/api/v1/settings/system", Method: model.EnumApiMethodGet, Group: "系统设置"},
		{Title: "设置系统设置", Path: "/api/v1/settings/system", Method: model.EnumApiMethodPut, Group: "系统设置"},
		{Title: "获取日志设置", Path: "/api/v1/settings/logger", Method: model.EnumApiMethodGet, Group: "系统设置"},
		{Title: "设置日志设置", Path: "/api/v1/settings/logger", Method: model.EnumApiMethodPut, Group: "系统设置"},

		// 字典管理
		{Title: "创建字典", Path: "/api/v1/dict_mana", Method: model.EnumApiMethodPost, Group: "字典管理"},
		{Title: "更新字典", Path: "/api/v1/dict_mana", Method: model.EnumApiMethodPut, Group: "字典管理"},
		{Title: "删除字典", Path: "/api/v1/dict_mana/{id}", Method: model.EnumApiMethodDelete, Group: "字典管理"},
		{Title: "查询字典列表", Path: "/api/v1/dict_mana/list", Method: model.EnumApiMethodPost, Group: "字典管理"},
		{Title: "根据id查询字典", Path: "/api/v1/dict_mana/{id}", Method: model.EnumApiMethodGet, Group: "字典管理"},
		{Title: "保存字典详情", Path: "/api/v1/dict_mana/detail", Method: model.EnumApiMethodPost, Group: "字典管理"},
		{Title: "查询字典详情列表", Path: "/api/v1/dict_mana/details", Method: model.EnumApiMethodPost, Group: "字典管理"},
		{Title: "删除字典详情", Path: "/api/v1/dict_mana/detail/{id}", Method: model.EnumApiMethodDelete, Group: "字典管理"},
		{Title: "查询字典详情", Path: "/api/v1/dict_mana/{id}/details", Method: model.EnumApiMethodGet, Group: "字典管理"},

		// 接口管理
	}

	lc := logic.NewApiManageLogic()
	for _, api := range apis {
		a := model.NewApi(api)
		slog.Info("检查接口数据", "title", api.Title, "path", api.Path)
		lc.IsNotExistAdd(a)
	}
}
