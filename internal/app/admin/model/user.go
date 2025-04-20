package model

import "lime/internal/common/model"

type User struct {
	model.BaseModel
	Username string `gorm:"column:username;unique;size:50;not null;comment:用户名，用于登录和识别用户;" json:"username"`
	Password string `gorm:"column:password;size:255;not null;comment:用户密码，以加密形式存储;" json:"password"`
	NickName string `gorm:"column:nick_name;size:50;not null;comment:用户昵称，用于显示在页面上，可修改;" json:"nick_name"`
	Avatar   string `gorm:"column:avatar;size:255;not null;comment:用户头像，用于显示在页面上，可修改;" json:"avatar"`
	Remark   string `gorm:"column:remark;size:255;not null;comment:用户备注，用于描述用户;" json:"remark"`
	// 一个操作员只能有一个角色
	UserRole *UserRole `gorm:"foreignKey:user_id"`
}

func (u User) TableName() string { return "sys_user" }

type Role struct {
	model.BaseModel
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

type UserRole struct {
	model.BaseModel
	UserID   uint   `gorm:"column:user_id;not null;comment:关联的用户ID;" json:"user_id"`
	RoleCode string `gorm:"column:role_code;not null;comment:关联的角色ID;" json:"role_code"`
}

func (ur UserRole) TableName() string { return "sys_user_role" }

type RoleMenu struct {
	model.BaseModel
	RoleCode string `gorm:"column:role_code;not null;comment:关联的角色编码;" json:"role_code"`
	MenuCode string `gorm:"column:menu_code;not null;comment:关联的菜单编码;" json:"menu_code"`
}

// TableName
func (RoleMenu) TableName() string { return "sys_role_menus" }
