package initialize

import (
	"lime/internal/app/admin/logic"
	"lime/internal/app/admin/model"
	"log/slog"
)

func InitApis() {
	apis := []model.ApiBase{
		// 基础接口
		{Title: "登录接口", Path: "/api/v1/auth/login", Method: model.HTTP_POST, ApiGroup: "基础接口"},
		{Title: "退出登录接口", Path: "/api/v1/auth/logout", Method: model.HTTP_POST, ApiGroup: "基础接口"},
		{Title: "刷新token接口", Path: "/api/v1/auth/refresh", Method: model.HTTP_GET, ApiGroup: "基础接口"},

		// 基础接口
		{Title: "获取用户信息接口", Path: "/api/v1/admin/userInfo", Method: model.HTTP_GET, ApiGroup: "基础接口"},
		{Title: "上传文件接口", Path: "/api/v1/admin/upload", Method: model.HTTP_POST, ApiGroup: "基础接口"},
		{Title: "更新用户信息接口", Path: "/api/v1/admin/updateuserinfo", Method: model.HTTP_POST, ApiGroup: "基础接口"},
		{Title: "更新密码接口", Path: "/api/v1/admin/updatepassword", Method: model.HTTP_POST, ApiGroup: "基础接口"},

		// 用户管理
		{Title: "根据条件查询数据", Path: "/api/v1/user/list", Method: model.HTTP_POST, ApiGroup: "用户管理"},
		{Title: "修改用户信息", Path: "/api/v1/user/update", Method: model.HTTP_PUT, ApiGroup: "用户管理"},
		{Title: "删除用户信息", Path: "/api/v1/user/delete/{id}", Method: model.HTTP_DELETE, ApiGroup: "用户管理"},
		{Title: "创建用户信息", Path: "/api/v1/user/add", Method: model.HTTP_POST, ApiGroup: "用户管理"},

		// 角色管理
		{Title: "根据条件查询数据", Path: "/api/v1/role/list", Method: model.HTTP_POST, ApiGroup: "角色管理"},
		{Title: "修改角色信息", Path: "/api/v1/role/update", Method: model.HTTP_PUT, ApiGroup: "角色管理"},
		{Title: "删除角色信息", Path: "/api/v1/role/delete/{id}", Method: model.HTTP_DELETE, ApiGroup: "角色管理"},
		{Title: "创建角色信息", Path: "/api/v1/role/add", Method: model.HTTP_POST, ApiGroup: "角色管理"},

		// 菜单管理
		{Title: "根据条件查询数据", Path: "/api/v1/menu/list", Method: model.HTTP_POST, ApiGroup: "菜单管理"},
		{Title: "获取所有菜单", Path: "/api/v1/menu/getAll", Method: model.HTTP_GET, ApiGroup: "菜单管理"},
		{Title: "获取角色菜单", Path: "/api/v1/menu/roleMenus/{code}", Method: model.HTTP_GET, ApiGroup: "菜单管理"},
		{Title: "保存角色菜单", Path: "/api/v1/menu/saveRoleMenus/{code}", Method: model.HTTP_POST, ApiGroup: "菜单管理"},
		{Title: "修改菜单", Path: "/api/v1/menu/update", Method: model.HTTP_PUT, ApiGroup: "菜单管理"},
		{Title: "删除菜单", Path: "/api/v1/menu/delete/{code}", Method: model.HTTP_PUT, ApiGroup: "菜单管理"},
		{Title: "创建菜单", Path: "/api/v1/menu/add", Method: model.HTTP_POST, ApiGroup: "菜单管理"},

		// 系统设置
		{Title: "获取系统设置", Path: "/api/v1/settings/system", Method: model.HTTP_GET, ApiGroup: "系统设置"},
		{Title: "设置系统设置", Path: "/api/v1/settings/system", Method: model.HTTP_PUT, ApiGroup: "系统设置"},
		{Title: "获取日志设置", Path: "/api/v1/settings/logger", Method: model.HTTP_GET, ApiGroup: "系统设置"},
		{Title: "设置日志设置", Path: "/api/v1/settings/logger", Method: model.HTTP_PUT, ApiGroup: "系统设置"},

		// 字典管理
		{Title: "创建字典", Path: "/api/v1/dict_mana", Method: model.HTTP_POST, ApiGroup: "字典管理"},
		{Title: "更新字典", Path: "/api/v1/dict_mana", Method: model.HTTP_PUT, ApiGroup: "字典管理"},
		{Title: "删除字典", Path: "/api/v1/dict_mana/{id}", Method: model.HTTP_DELETE, ApiGroup: "字典管理"},
		{Title: "查询字典列表", Path: "/api/v1/dict_mana/list", Method: model.HTTP_POST, ApiGroup: "字典管理"},
		{Title: "根据id查询字典", Path: "/api/v1/dict_mana/{id}", Method: model.HTTP_GET, ApiGroup: "字典管理"},
		{Title: "保存字典详情", Path: "/api/v1/dict_mana/detail", Method: model.HTTP_POST, ApiGroup: "字典管理"},
		{Title: "查询字典详情列表", Path: "/api/v1/dict_mana/details", Method: model.HTTP_POST, ApiGroup: "字典管理"},
		{Title: "删除字典详情", Path: "/api/v1/dict_mana/detail/{id}", Method: model.HTTP_DELETE, ApiGroup: "字典管理"},
		{Title: "查询字典详情", Path: "/api/v1/dict_mana/{id}/details", Method: model.HTTP_GET, ApiGroup: "字典管理"},

		// 接口管理
	}

	lc := logic.NewApiManageLogic()
	for _, api := range apis {
		a := model.NewApi(api)
		slog.Info("检查接口数据", "title", api.Title, "path", api.Path)
		lc.IsNotExistAdd(a)
	}
}
