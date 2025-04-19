package route

import (
	v1 "lime/internal/app/admin/controller/v1"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	auth := r.Group("auth")
	{
		auth.POST("login", v1.Login)         // 登录接口
		auth.POST("logout", v1.Logout)       // 退出登录接口
		auth.GET("refresh", v1.RefreshToken) // 刷新token接口
	}

	admin := r.Group("admin")
	{
		admin.GET("userInfo", v1.GetUserInfo)           // 获取用户信息接口
		admin.POST("upload", v1.Upload)                 // 上传文件接口
		admin.POST("updateuserinfo", v1.UpdateUserinfo) // 更新用户信息接口
		admin.POST("updatepassword", v1.UpdatePassword) // 更新密码接口
		admin.GET("homeCount", v1.GetHomeCount)         // 获取首页统计信息接口
	}

	user := r.Group("user")
	{
		user.POST("list", v1.GetUsers)           // 根据条件查询数据
		user.PUT("update", v1.UpdateUser)        // 修改用户信息
		user.DELETE("delete/:id", v1.DeleteUser) // 删除用户信息
		user.POST("add", v1.CreateUser)          // 创建用户信息
	}

	role := r.Group("role")
	{
		role.POST("list", v1.GetRoles)           // 根据条件查询数据
		role.PUT("update", v1.UpdateRole)        // 修改角色信息
		role.DELETE("delete/:id", v1.DeleteRole) // 删除角色信息
		role.POST("add", v1.CreateRole)          // 创建角色信息
	}

	menu := r.Group("menu")
	{
		menu.POST("list", v1.GetMenus)                     // 根据条件查询数据
		menu.GET("getAll", v1.GetAll)                      // 获取所有菜单
		menu.GET("roleMenus/:code", v1.GetRoleMenus)       // 获取角色菜单
		menu.POST("saveRoleMenus/:code", v1.SaveRoleMenus) // 保存角色菜单
		menu.PUT("update", v1.UpdateMenu)                  // 修改菜单
		menu.DELETE("delete/:id", v1.DeleteMenu)           // 删除菜单
		menu.POST("add", v1.CreateMenu)                    // 创建菜单
	}

	settings := r.Group("settings")
	{
		settings.GET("system", v1.GetSystemSettings) // 获取系统设置
		settings.PUT("system", v1.SetSystemSettings) // 设置系统设置
		settings.GET("logger", v1.GetLoggerSettings) // 获取日志设置
		settings.PUT("logger", v1.SetLoggerSettings) // 设置日志设置
	}

	dict_mana := r.Group("dict_mana")
	{
		dict_mana.POST("", v1.CreateDicts)              // 创建字典
		dict_mana.PUT("", v1.UpdateDicts)               // 更新字典
		dict_mana.DELETE(":id", v1.DeleteDicts)         // 删除字典
		dict_mana.POST("list", v1.DictsList)            // 查询字典列表
		dict_mana.GET(":id", v1.GetDicts)               // 查询字典
		dict_mana.POST("detail", v1.SaveDetail)         // 保存字典详情
		dict_mana.POST("details", v1.ListDetail)        // 查询字典详情列表
		dict_mana.DELETE("detail/:id", v1.DelDetail)    // 删除字典详情
		dict_mana.GET(":id/details", v1.GetDictDetails) // 查询字典详情
	}
}
