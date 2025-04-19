package requests

import (
	"lime/internal/app/admin/model"
	commonModel "lime/internal/common/model"
)

type QueryMenu struct {
	KeyWords string `json:"keywords" form:"keywords"`   // 关键词
	IsHidden int    `json:"is_hidden" form:"is_hidden"` // 0 不隐藏 1 隐藏
}

func NewQueryMenu() *commonModel.PageQuery[*QueryMenu] {
	return commonModel.NewPageQuery(0, 0, &QueryMenu{})
}

type CreateMenu struct {
	Code        string `json:"code" binding:"required"` // 标识码
	Name        string `json:"name" binding:"required"` // 标题
	Description string `json:"desc"`                    // 描述
	Frontpath   string `json:"frontpath"`               // 前端路径
	Order       int    `json:"order"`                   // 排序
	Icon        string `json:"icon"`                    // 图标
	ParentCode  string `json:"parent_code"`             // 父级菜单标识码
	MenuType    int    `json:"menu_type"`               // 菜单类型
	IsLink      int    `json:"is_link"`                 // 是否外链
}

func NewCreateMenu() *CreateMenu {
	return &CreateMenu{}
}

type UpdateMenu struct {
	Id          int                `json:"id" binding:"required"`   // 编码
	Code        string             `json:"code" binding:"required"` // 标识码
	Name        string             `json:"name" binding:"required"` // 标题
	Description string             `json:"desc"`                    // 描述
	Frontpath   string             `json:"frontpath"`               // 前端路径
	Order       int                `json:"order"`                   // 排序
	Icon        string             `json:"icon"`                    // 图标
	ParentCode  string             `json:"parent_code"`             // 父级菜单标识码
	MenuType    model.EnumMenuType `json:"menu_type"`               // 菜单类型
	IsLink      int                `json:"is_link"`                 // 是否外链
}

func NewUpdateMenu() *UpdateMenu {
	return &UpdateMenu{}
}
