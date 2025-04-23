package model

import (
	"lime/internal/common"
	"lime/internal/common/model"
)

type ApiInfo struct {
	model.BaseModel
	ApiBase
}

type ApiBase struct {
	Title    string               `gorm:"column:title;size:50;not null;comment:接口标题;" json:"title"`     // 接口标题
	Path     string               `gorm:"column:path;size:255;not null;comment:接口路径;" json:"path"`      // 接口路径
	Method   common.EnumApiMethod `gorm:"column:method;size:10;not null;comment:接口方法;" json:"method"`   // 接口方法
	ApiGroup string               `gorm:"column:api_group;size:50;not null;comment:接口分组;" json:"group"` // 接口分组
}

func (d ApiInfo) TableName() string { return "sys_api" }

func NewApi(data ApiBase) *ApiInfo {
	return &ApiInfo{
		ApiBase: data,
	}
}
