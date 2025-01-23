package response

import "lime/internal/app/admin/model"

type Auth struct {
	Token string `json:"token"`
	User  string `json:"user"`
	ID    uint   `json:"id"`
}

type UserInfo struct {
	model.User
	Menus []*model.Menu `json:"menus"`
}

type RoleMenuResponse struct {
	model.Menu
	IsHave bool `gorm:"column:is_have;comment:是否可见;" json:"is_have"` // 角色是否有菜单权限
}

type Settings struct {
	Group string `json:"group"`
	Key   string `json:"key"`
	Value string `json:"value" binding:"required"`
}
