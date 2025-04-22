package route

import (
	v1 "lime/internal/app/admin/controller/v1"

	"github.com/gin-gonic/gin"
)

type Router struct {
	ApiHandlers     v1.ApiManageController
	UserHandlers    v1.UserController
	RoleHandlers    v1.RoleController
	MenuHandlers    v1.MenuController
	DictHandlers    v1.DictController
	SettingHandlers v1.SettingsController
}

func MakeRouter() Router {
	return Router{
		ApiHandlers:     v1.MakeApiManageController(),
		UserHandlers:    v1.MakeUserController(),
		RoleHandlers:    v1.MakeRoleController(),
		MenuHandlers:    v1.MakeMenuController(),
		DictHandlers:    v1.MakeDictController(),
		SettingHandlers: v1.MakeSettingsController(),
	}
}

func (router *Router) Register(r *gin.RouterGroup) {
	auth := r.Group("auth")
	{
		auth.POST("login", v1.Login)         // 登录接口
		auth.POST("logout", v1.Logout)       // 退出登录接口
		auth.GET("refresh", v1.RefreshToken) // 刷新token接口
	}

	admin := r.Group("admin")
	{
		admin.GET("userInfo", router.UserHandlers.GetUserInfo)           // 获取用户信息接口
		admin.POST("upload", v1.Upload)                                  // 上传文件接口
		admin.POST("updateuserinfo", router.UserHandlers.UpdateUserinfo) // 更新用户信息接口
		admin.POST("updatepassword", router.UserHandlers.UpdatePassword) // 更新密码接口
	}

	user := r.Group("user")
	{
		user.POST("list", router.UserHandlers.GetUsers)           // 根据条件查询数据
		user.PUT("update", router.UserHandlers.UpdateUser)        // 修改用户信息
		user.DELETE("delete/:id", router.UserHandlers.DeleteUser) // 删除用户信息
		user.POST("add", router.UserHandlers.CreateUser)          // 创建用户信息
	}

	role := r.Group("role")
	{
		role.POST("list", router.RoleHandlers.Roles)          // 根据条件查询数据
		role.PUT("update", router.RoleHandlers.Update)        // 修改角色信息
		role.DELETE("delete/:id", router.RoleHandlers.Delete) // 删除角色信息
		role.POST("add", router.RoleHandlers.Create)          // 创建角色信息
	}

	menu := r.Group("menu")
	{
		menu.POST("list", router.MenuHandlers.GetMenus)                     // 根据条件查询数据
		menu.GET("getAll", router.MenuHandlers.GetAll)                      // 获取所有菜单
		menu.GET("roleMenus/:code", router.MenuHandlers.GetRoleMenus)       // 获取角色菜单
		menu.POST("saveRoleMenus/:code", router.MenuHandlers.SaveRoleMenus) // 保存角色菜单
		menu.PUT("update", router.MenuHandlers.Update)                      // 修改菜单
		menu.DELETE("delete/:id", router.MenuHandlers.Delete)               // 删除菜单
		menu.POST("add", router.MenuHandlers.Create)                        // 创建菜单
	}

	settings := r.Group("settings")
	{
		settings.GET("system", router.SettingHandlers.GetSystemSettings) // 获取系统设置
		settings.PUT("system", router.SettingHandlers.SetSystemSettings) // 设置系统设置
		settings.GET("logger", router.SettingHandlers.GetLoggerSettings) // 获取日志设置
		settings.PUT("logger", router.SettingHandlers.SetLoggerSettings) // 设置日志设置
	}

	dict_mana := r.Group("dict_mana")
	{
		dict_mana.POST("", router.DictHandlers.CreateDicts)              // 创建字典
		dict_mana.PUT("", router.DictHandlers.UpdateDicts)               // 更新字典
		dict_mana.DELETE(":id", router.DictHandlers.DeleteDicts)         // 删除字典
		dict_mana.POST("list", router.DictHandlers.DictsList)            // 查询字典列表
		dict_mana.GET(":id", router.DictHandlers.GetDicts)               // 查询字典
		dict_mana.POST("detail", router.DictHandlers.SaveDetail)         // 保存字典详情
		dict_mana.POST("details", router.DictHandlers.ListDetail)        // 查询字典详情列表
		dict_mana.DELETE("detail/:id", router.DictHandlers.DelDetail)    // 删除字典详情
		dict_mana.GET(":id/details", router.DictHandlers.GetDictDetails) // 查询字典详情
	}

	api_manage := r.Group("api_manage")
	{
		api_manage.POST("", router.ApiHandlers.CreateApiInfo)      // 创建接口
		api_manage.PUT("", router.ApiHandlers.UpdateApiInfo)       // 更新接口
		api_manage.DELETE(":id", router.ApiHandlers.DeleteApiInfo) // 删除接口
		api_manage.POST("list", router.ApiHandlers.GetApiInfos)    // 查询接口列表
	}
}
