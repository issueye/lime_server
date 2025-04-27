package model

import "lime/internal/common/model"

type RoleMenu struct {
	model.BaseModel
	RoleCode string `gorm:"column:role_code;not null;comment:关联的角色编码;" json:"role_code"`
	MenuCode string `gorm:"column:menu_code;not null;comment:关联的菜单编码;" json:"menu_code"`
}

// TableName
func (RoleMenu) TableName() string { return "sys_role_menus" }
