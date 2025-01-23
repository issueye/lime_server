package route

import (
	v1 "lime/internal/app/admin/controller/v1"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	auth := r.Group("auth")
	{
		auth.POST("login", v1.Login)
		auth.POST("logout", v1.Logout)
		auth.GET("refresh", v1.RefreshToken)
	}

	admin := r.Group("admin")
	{
		admin.GET("userInfo", v1.GetUserInfo)
		admin.POST("upload", v1.Upload)
		admin.POST("updateuserinfo", v1.UpdateUserinfo)
		admin.POST("updatepassword", v1.UpdatePassword)
		admin.GET("homeCount", v1.GetHomeCount)
	}

	user := r.Group("user")
	{
		user.POST("list", v1.GetUsers)
		user.PUT("update", v1.UpdateUser)
		user.DELETE("delete/:id", v1.DeleteUser)
		user.POST("add", v1.CreateUser)
	}

	role := r.Group("role")
	{
		role.POST("list", v1.GetRoles)
		role.PUT("update", v1.UpdateRole)
		role.DELETE("delete/:id", v1.DeleteRole)
		role.POST("add", v1.CreateRole)
	}

	menu := r.Group("menu")
	{
		menu.POST("list", v1.GetMenus)
		menu.GET("roleMenus/:code", v1.GetRoleMenus)
		menu.GET("catalog", v1.GetCatalog)
		menu.POST("saveRoleMenus/:code", v1.SaveRoleMenus)
		menu.PUT("update", v1.UpdateMenu)
		menu.DELETE("delete/:id", v1.DeleteMenu)
		menu.POST("add", v1.CreateMenu)
	}

	settings := r.Group("settings")
	{
		settings.GET("system", v1.GetSystemSettings)
		settings.PUT("system", v1.SetSystemSettings)
		settings.GET("logger", v1.GetLoggerSettings)
		settings.PUT("logger", v1.SetLoggerSettings)
	}

	dict_mana := r.Group("dict_mana")
	{
		dict_mana.POST("", v1.CreateDicts)
		dict_mana.PUT("", v1.UpdateDicts)
		dict_mana.DELETE(":id", v1.DeleteDicts)
		dict_mana.POST("list", v1.DictsList)
		dict_mana.GET(":id", v1.GetDicts)
		dict_mana.POST("detail", v1.SaveDetail)
		dict_mana.POST("details", v1.ListDetail)
		dict_mana.DELETE("detail/:id", v1.DelDetail)
		dict_mana.GET(":id/details", v1.GetDictDetails)
	}
}
