package requests

import (
	commonModel "lime/internal/common/model"
)

type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type QueryUser struct {
	KeyWords string `json:"keywords" form:"keywords"`
}

func NewQueryUser() *commonModel.PageQuery[*QueryUser] {
	return commonModel.NewPageQuery(0, 0, &QueryUser{})
}

type UpdatePassword struct {
	Oldpassword string `json:"oldpassword"`
	Password    string `json:"password"`
	Repassword  string `json:"repassword"`
}

func NewUpdatePassword() *UpdatePassword {
	return &UpdatePassword{}
}

type UpdateUser struct {
	Id       int    `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	NickName string `json:"nick_name"  binding:"required"`
	Avatar   string `json:"avatar"`
}

func NewUpdateUser() *UpdateUser {
	return &UpdateUser{}
}

type CreateUser struct {
	Username string `json:"username" binding:"required"`
	NickName string `json:"nick_name"  binding:"required"`
	RoleCode string `json:"role_code" binding:"required"`
	Avatar   string `json:"avatar"`
}

func NewCreateUser() *CreateUser {
	return &CreateUser{}
}

type QueryRole struct {
	KeyWords string `json:"keywords" form:"keywords"`
}

func NewQueryRole() *commonModel.PageQuery[*QueryRole] {
	return commonModel.NewPageQuery(0, 0, &QueryRole{})
}

type CreateRole struct {
	Code   string `json:"code" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Remark string `json:"remark"`
}

func NewCreateRole() *CreateRole {
	return &CreateRole{}
}

type UpdateRole struct {
	Id        int      `json:"id" binding:"required"`
	Code      string   `json:"code" binding:"required"`
	Name      string   `json:"name" binding:"required"`
	Remark    string   `json:"remark"`
	MenuCodes []string `json:"menu_codes"`
}

func NewUpdateRole() *UpdateRole {
	return &UpdateRole{}
}

type Settings struct {
	Group string `json:"group" binding:"required"`
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

func NewSettings() []*Settings {
	return []*Settings{}
}

type RoleApi struct {
	RoleCode string `json:"role_code" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Method   string `json:"method" binding:"required"`
}

func NewRoleApi() *RoleApi {
	return &RoleApi{}
}

type CreateRoleApis struct {
	RoleCode string `json:"role_code" binding:"required"` // 角色编码
	Apis     []uint `json:"apis" binding:"required"`      // 接口id列表
}

func NewCreateRoleApis() *CreateRoleApis {
	return &CreateRoleApis{}
}

type RoleQryApi struct {
	RoleCode string `json:"role_code" label:"角色编码" binding:"required"`
	Keyword  string `json:"keyword" label:"关键词" form:"keyword"`
}

func NewRoleQryApi() *RoleQryApi {
	return &RoleQryApi{}
}
