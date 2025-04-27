package initialize

import (
	"context"
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/service"
	"lime/internal/common"
	"lime/internal/global"
	"log/slog"

	adapter "github.com/casbin/gorm-adapter/v3"
)

// 初始化角色数据
func InitRoles() {
	Roles := []model.RoleBase{
		{Code: "9999", Name: "管理员", IsCanNotRemove: 1, Remark: "系统生成"},
	}

	for _, Role := range Roles {
		data := model.NewRole(Role)
		slog.Info("检查角色数据", "code", Role.Code, "name", Role.Name)
		RoleIsNotExistAdd(data)
	}
}

func RoleIsNotExistAdd(Role *model.Role) {
	roleSrv := service.NewRole()
	info, err := roleSrv.GetByField("code", Role.Code)
	if err != nil {
		slog.Error("查询角色信息失败", slog.String("错误原因", err.Error()))
		return
	}

	if info.ID == 0 {
		err = roleSrv.Create(Role)
		if err != nil {
			slog.Error("添加角色失败", slog.String("错误原因", err.Error()))
			return
		}
	}
}

// 初始化角色与用户关联
func InitUserRole() {
	userRole := []*model.UserRole{
		{UserID: 1, RoleCode: "9999"},
	}

	for _, ur := range userRole {
		URIsNotExistAdd(ur)
	}
}

func URIsNotExistAdd(ur *model.UserRole) {
	RoleSrv := service.NewUser()
	isHave, err := RoleSrv.CheckUserRole(int(ur.UserID), ur.RoleCode)
	if err != nil {
		slog.Error("查询用户角色失败", slog.String("错误原因", err.Error()))
		return
	}

	if !isHave {
		err = RoleSrv.AddUserRole(ur)
		if err != nil {
			slog.Error("添加用户角色失败", slog.String("错误原因", err.Error()))
			return
		}
	}
}

// 初始化角色与菜单关联
func InitRoleMenus() {
	rms := []*model.RoleMenu{
		{RoleCode: "9999", MenuCode: "9000"},
		{RoleCode: "9999", MenuCode: "9001"},
		{RoleCode: "9999", MenuCode: "9002"},
		{RoleCode: "9999", MenuCode: "9003"},
		{RoleCode: "9999", MenuCode: "9004"},
	}

	for _, rm := range rms {
		RMIsNotExistAdd(rm)
	}
}

func RMIsNotExistAdd(rm *model.RoleMenu) {
	RoleSrv := service.NewUser()
	isHave, err := RoleSrv.CheckRoleMenu(rm.RoleCode, rm.MenuCode)
	if err != nil {
		slog.Error("查询角色菜单失败", slog.String("错误原因", err.Error()))
		return
	}

	if !isHave {
		err = RoleSrv.AddRoleMenu(rm)
		if err != nil {
			slog.Error("添加角色菜单失败，失败原因：%s", slog.String("错误原因", err.Error()))
			return
		}
	}
}

func InitRoleApi() {
	apis := []adapter.CasbinRule{
		// 基础接口
		{Ptype: "p", V0: "9999", V1: "/auth/login", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/auth/logout", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/auth/refresh", V2: common.HTTP_GET.String(), V3: "", V4: "", V5: ""},
		// 用户信息
		{Ptype: "p", V0: "9999", V1: "/admin/userInfo", V2: common.HTTP_GET.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/admin/upload", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/admin/updateuserinfo", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/admin/updatepassword", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		// 用户管理
		{Ptype: "p", V0: "9999", V1: "/user/list", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/user/update", V2: common.HTTP_PUT.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/user/delete/:id", V2: common.HTTP_DELETE.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/user/add", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/user/resetPwd/:id", V2: common.HTTP_GET.String(), V3: "", V4: "", V5: ""},
		// 角色管理
		{Ptype: "p", V0: "9999", V1: "/role/list", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/role/update", V2: common.HTTP_PUT.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/role/delete/:id", V2: common.HTTP_DELETE.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/role/add", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/role/getApis", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/role/getNoHaveApis", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/role/removeApi", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/role/addApi", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/role/freshCasbin", V2: common.HTTP_GET.String(), V3: "", V4: "", V5: ""},
		// 菜单管理
		{Ptype: "p", V0: "9999", V1: "/menu/list", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/menu/getAll", V2: common.HTTP_GET.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/menu/roleMenus/:code", V2: common.HTTP_GET.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/menu/saveRoleMenus/:code", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/menu/update", V2: common.HTTP_PUT.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/menu/delete/:code", V2: common.HTTP_DELETE.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/menu/add", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		// 系统设置
		{Ptype: "p", V0: "9999", V1: "/settings/system", V2: common.HTTP_GET.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/settings/system", V2: common.HTTP_PUT.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/settings/logger", V2: common.HTTP_GET.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/settings/logger", V2: common.HTTP_PUT.String(), V3: "", V4: "", V5: ""},
		// 字典管理
		{Ptype: "p", V0: "9999", V1: "/dict_mana", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/dict_mana", V2: common.HTTP_PUT.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/dict_mana/list", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/dict_mana/:id", V2: common.HTTP_DELETE.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/dict_mana/detail", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/dict_mana/details", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/dict_mana/detail/:id", V2: common.HTTP_DELETE.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/dict_mana/:id/details", V2: common.HTTP_GET.String(), V3: "", V4: "", V5: ""},
		// 接口管理
		{Ptype: "p", V0: "9999", V1: "/api_manage", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/api_manage", V2: common.HTTP_PUT.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/api_manage/list", V2: common.HTTP_POST.String(), V3: "", V4: "", V5: ""},
		{Ptype: "p", V0: "9999", V1: "/api_manage/:id", V2: common.HTTP_DELETE.String(), V3: "", V4: "", V5: ""},
	}

	for _, api := range apis {
		ApiIsNotExistAdd(&api)
	}
}

func ApiIsNotExistAdd(api *adapter.CasbinRule) {
	db := global.DB.WithContext(context.Background())

	// 查询是否存在
	var count int64
	err := db.Model(&adapter.CasbinRule{}).Where("v0 = ? AND v1 = ? AND v2 = ?", api.V0, api.V1, api.V2).Count(&count).Error
	if err != nil {
		slog.Error("查询角色菜单失败", slog.String("错误原因", err.Error()))
		return
	}

	if count == 0 {
		err = db.Create(api).Error
		if err != nil {
			slog.Error("添加角色菜单失败", slog.String("错误原因", err.Error()))
			return
		}
	}
}
