package model

import "lime/internal/common/model"

type Role struct {
	model.BaseModel
	// 角色菜单关系表 关联 sys_role_menus 表 角色 code 关联菜单 code
	RoleMenus []*RoleMenu `gorm:"foreignKey:role_code;references:code;" json:"role_menus"`
	RoleBase
}

type RoleBase struct {
	Code           string `gorm:"column:code;unique;size:50;not null;comment:角色编码，用于标识角色;" json:"code"`     // 角色编码
	Name           string `gorm:"column:name;unique;size:50;not null;comment:角色名称，如管理员、普通用户等;" json:"name"` // 角色名称
	IsCanNotRemove uint   `gorm:"column:is_can_not_remove;comment:是否可删除;" json:"is_can_not_remove"`         // 是否可删除 0 否 1 是
	Remark         string `gorm:"column:remark;size:255;not null;comment:角色备注，用于描述角色;" json:"remark"`       // 角色备注
}

func (r Role) TableName() string {
	return "sys_role"
}

func NewRole(data RoleBase) *Role {
	return &Role{
		RoleBase: data,
	}
}
