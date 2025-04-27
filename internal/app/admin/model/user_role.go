package model

import "lime/internal/common/model"

type UserRole struct {
	model.BaseModel
	UserID   uint   `gorm:"column:user_id;not null;comment:关联的用户ID;" json:"user_id"`
	RoleCode string `gorm:"column:role_code;not null;comment:关联的角色ID;" json:"role_code"`
}

func (ur UserRole) TableName() string { return "sys_user_role" }
