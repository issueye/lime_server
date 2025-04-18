package model

import "lime/internal/common/model"

type Menu struct {
	model.BaseModel
	MenuBase
	IsHave   bool    `gorm:"column:is_have;comment:是否可见;" json:"is_have"` // 是否有权限
	Children []*Menu `gorm:"-" json:"child"`
}

type EnumMenuType uint

const (
	EMT_MENU      EnumMenuType = 0
	EMT_DIRECTORY EnumMenuType = 1
)

type MenuBase struct {
	Code        string       `gorm:"column:code;size:200;not null;comment:菜单编码;" json:"code"`        // 菜单编码
	Name        string       `gorm:"column:name;size:200;not null;comment:菜单名称;" json:"name"`        // 菜单名称
	Description string       `gorm:"column:description;size:200;comment:菜单描述;" json:"desc"`          // 菜单描述
	Frontpath   string       `gorm:"column:frontpath;size:200;comment:前端路由地址;" json:"frontpath"`     // 前端路由地址
	Order       int          `gorm:"column:order;comment:菜单排序;" json:"order"`                        // 菜单排序
	Icon        string       `gorm:"column:icon;size:200;comment:菜单图标;" json:"icon"`                 // 菜单图标
	Visible     bool         `gorm:"column:visible;comment:是否可见;" json:"visible"`                    // 是否可见
	ParentCode  string       `gorm:"column:parent_code;size:200;comment:父级菜单编码;" json:"parent_code"` // 父级菜单编码
	MenuType    EnumMenuType `gorm:"column:menu_type;comment:菜单类型;" json:"menu_type"`                // 菜单类型
}

func BaseNewMenu(base MenuBase) *Menu {
	return &Menu{
		MenuBase: base,
	}
}

func (Menu) TableName() string { return "sys_menu" }

func (m *Menu) GetCode() string {
	return m.Code
}

func (m *Menu) GetParentCode() string {
	return m.ParentCode
}

func (m *Menu) GetChildren() []*Menu {
	return m.Children
}
