package requests

import (
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
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Frontpath   string `json:"frontpath"`
	Condition   string `json:"condition"`
	Order       int    `json:"order"`
	Icon        string `json:"icon"`
	Method      string `json:"method"`
	ParentCode  string `json:"parent_code"`
}

func NewCreateMenu() *CreateMenu {
	return &CreateMenu{}
}

type UpdateMenu struct {
	Id          int    `json:"id" binding:"required"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Frontpath   string `json:"frontpath"`
	Condition   string `json:"condition"`
	Order       int    `json:"order"`
	Icon        string `json:"icon"`
	Method      string `json:"method"`
	ParentCode  string `json:"parent_code"`
}

func NewUpdateMenu() *UpdateMenu {
	return &UpdateMenu{}
}
