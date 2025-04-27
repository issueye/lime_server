package initialize

import (
	"lime/internal/app/admin/logic"
	"lime/internal/app/admin/model"
	"lime/internal/common"
	"log/slog"
)

func InitApis() {
	apis := []model.ApiBase{
		// 基础接口
		{Title: "登录接口", Path: "/api/v1/auth/login", Method: common.HTTP_POST, ApiGroup: "基础接口"},
		{Title: "退出登录接口", Path: "/api/v1/auth/logout", Method: common.HTTP_POST, ApiGroup: "基础接口"},
		{Title: "刷新token接口", Path: "/api/v1/auth/refresh", Method: common.HTTP_GET, ApiGroup: "基础接口"},

		// 基础接口
		{Title: "获取用户信息接口", Path: "/api/v1/admin/userInfo", Method: common.HTTP_GET, ApiGroup: "基础接口"},
		{Title: "上传文件接口", Path: "/api/v1/admin/upload", Method: common.HTTP_POST, ApiGroup: "基础接口"},
		{Title: "更新用户信息接口", Path: "/api/v1/admin/updateuserinfo", Method: common.HTTP_POST, ApiGroup: "基础接口"},
		{Title: "更新密码接口", Path: "/api/v1/admin/updatepassword", Method: common.HTTP_POST, ApiGroup: "基础接口"},

		// 用户管理
		{Title: "根据条件查询数据", Path: "/api/v1/user/list", Method: common.HTTP_POST, ApiGroup: "用户管理"},
		{Title: "修改用户信息", Path: "/api/v1/user/update", Method: common.HTTP_PUT, ApiGroup: "用户管理"},
		{Title: "删除用户信息", Path: "/api/v1/user/delete/:id", Method: common.HTTP_DELETE, ApiGroup: "用户管理"},
		{Title: "创建用户信息", Path: "/api/v1/user/add", Method: common.HTTP_POST, ApiGroup: "用户管理"},
		{Title: "重置用户密码", Path: "/api/v1/user/resetPwd/:id", Method: common.HTTP_PUT, ApiGroup: "用户管理"},

		// 角色管理
		{Title: "根据条件查询数据", Path: "/api/v1/role/list", Method: common.HTTP_POST, ApiGroup: "角色管理"},
		{Title: "修改角色信息", Path: "/api/v1/role/update", Method: common.HTTP_PUT, ApiGroup: "角色管理"},
		{Title: "删除角色信息", Path: "/api/v1/role/delete/:id", Method: common.HTTP_DELETE, ApiGroup: "角色管理"},
		{Title: "创建角色信息", Path: "/api/v1/role/add", Method: common.HTTP_POST, ApiGroup: "角色管理"},

		// 菜单管理
		{Title: "根据条件查询数据", Path: "/api/v1/menu/list", Method: common.HTTP_POST, ApiGroup: "菜单管理"},
		{Title: "获取所有菜单", Path: "/api/v1/menu/getAll", Method: common.HTTP_GET, ApiGroup: "菜单管理"},
		{Title: "获取角色菜单", Path: "/api/v1/menu/roleMenus/:code", Method: common.HTTP_GET, ApiGroup: "菜单管理"},
		{Title: "保存角色菜单", Path: "/api/v1/menu/saveRoleMenus/:code", Method: common.HTTP_POST, ApiGroup: "菜单管理"},
		{Title: "修改菜单", Path: "/api/v1/menu/update", Method: common.HTTP_PUT, ApiGroup: "菜单管理"},
		{Title: "删除菜单", Path: "/api/v1/menu/delete/:code", Method: common.HTTP_PUT, ApiGroup: "菜单管理"},
		{Title: "创建菜单", Path: "/api/v1/menu/add", Method: common.HTTP_POST, ApiGroup: "菜单管理"},

		// 系统设置
		{Title: "获取系统设置", Path: "/api/v1/settings/system", Method: common.HTTP_GET, ApiGroup: "系统设置"},
		{Title: "设置系统设置", Path: "/api/v1/settings/system", Method: common.HTTP_PUT, ApiGroup: "系统设置"},
		{Title: "获取日志设置", Path: "/api/v1/settings/logger", Method: common.HTTP_GET, ApiGroup: "系统设置"},
		{Title: "设置日志设置", Path: "/api/v1/settings/logger", Method: common.HTTP_PUT, ApiGroup: "系统设置"},

		// 字典管理
		{Title: "创建字典", Path: "/api/v1/dict_mana", Method: common.HTTP_POST, ApiGroup: "字典管理"},
		{Title: "更新字典", Path: "/api/v1/dict_mana", Method: common.HTTP_PUT, ApiGroup: "字典管理"},
		{Title: "删除字典", Path: "/api/v1/dict_mana/:id", Method: common.HTTP_DELETE, ApiGroup: "字典管理"},
		{Title: "查询字典列表", Path: "/api/v1/dict_mana/list", Method: common.HTTP_POST, ApiGroup: "字典管理"},
		{Title: "根据id查询字典", Path: "/api/v1/dict_mana/:id", Method: common.HTTP_GET, ApiGroup: "字典管理"},
		{Title: "保存字典详情", Path: "/api/v1/dict_mana/detail", Method: common.HTTP_POST, ApiGroup: "字典管理"},
		{Title: "查询字典详情列表", Path: "/api/v1/dict_mana/details", Method: common.HTTP_POST, ApiGroup: "字典管理"},
		{Title: "删除字典详情", Path: "/api/v1/dict_mana/detail/:id", Method: common.HTTP_DELETE, ApiGroup: "字典管理"},
		{Title: "查询字典详情", Path: "/api/v1/dict_mana/:id/details", Method: common.HTTP_GET, ApiGroup: "字典管理"},

		// 接口管理
		{Title: "查询接口信息列表", Path: "/api/v1/api_manage/list", Method: common.HTTP_POST, ApiGroup: "接口管理"},
		{Title: "创建接口", Path: "/api/v1/api_manage", Method: common.HTTP_POST, ApiGroup: "接口管理"},
		{Title: "更新接口", Path: "/api/v1/api_manage", Method: common.HTTP_PUT, ApiGroup: "接口管理"},
		{Title: "删除", Path: "/api/v1/api_manage/:id", Method: common.HTTP_DELETE, ApiGroup: "接口管理"},
	}

	lc := logic.NewApiManageLogic()
	for _, api := range apis {
		a := model.NewApi(api)
		slog.Info("检查接口数据", "title", api.Title, "path", api.Path)
		lc.IsNotExistAdd(a)
	}
}
