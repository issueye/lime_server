package requests

import (
	"lime/internal/common"
	commonModel "lime/internal/common/model"
)

type QueryApiInfo struct {
	KeyWords string `json:"keywords" form:"keywords"` // 关键词
}

func NewQueryApiInfo() *commonModel.PageQuery[*QueryApiInfo] {
	return commonModel.NewPageQuery(0, 0, &QueryApiInfo{})
}

type CreateApiInfo struct {
	Title  string               `json:"title" binding:"required"`  // 接口标题
	Path   string               `json:"path" binding:"required"`   // 接口路径
	Method common.EnumApiMethod `json:"method" binding:"required"` // 接口方法
	Group  string               `json:"group" binding:"required"`  // 接口分组
}

func NewCreateApiInfo() *CreateApiInfo {
	return &CreateApiInfo{}
}

type UpdateApiInfo struct {
	Id     uint                 `json:"id" binding:"required"`     // 编码
	Title  string               `json:"title" binding:"required"`  // 接口标题
	Path   string               `json:"path" binding:"required"`   // 接口路径
	Method common.EnumApiMethod `json:"method" binding:"required"` // 接口方法
	Group  string               `json:"group"`                     // 接口分组
}

func NewUpdateApiInfo() *UpdateApiInfo {
	return &UpdateApiInfo{}
}
