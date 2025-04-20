package initialize

import (
	"lime/internal/app/admin/logic"
	"lime/internal/app/admin/model"
	"log/slog"
)

func InitAdminData() {
	InitRoles()
	InitMenus()
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
