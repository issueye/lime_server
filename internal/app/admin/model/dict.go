package model

import (
	"lime/internal/common"
	"lime/internal/common/model"
)

type DictsInfo struct {
	model.BaseModel
	DictsBase
}

func NewDictsInfo(base DictsBase) *DictsInfo {
	srv := &DictsInfo{
		DictsBase: base,
	}
	return srv
}

type DictsBase struct {
	Code        string                     `gorm:"column:code;size:200;not null;comment:编码;" json:"code"`               // 编码
	Name        string                     `gorm:"column:name;size:200;not null;comment:名称;" json:"name"`               // 名称
	ContentType common.EnumDictContentType `gorm:"column:content_type;not null;comment:内容类型;" json:"content_type"`      // 内容类型
	Description string                     `gorm:"column:description;size:255;not null;comment:描述;" json:"description"` // 描述
	Details     []DictDetail               `gorm:"foreignKey:DictCode;references:code;" json:"details"`                 // 字典详情
}

func (d DictsInfo) TableName() string { return "sys_dict_info" }

type DictDetail struct {
	DictCode    string `gorm:"column:dict_code;size:200;not null;comment:字典编码;" json:"dict_code"`   // 字典编码
	Key         string `gorm:"column:key;size:200;not null;comment:字典键;" json:"key"`                // 字典键
	Value       string `gorm:"column:value;size:200;not null;comment:字典值;" json:"val"`              // 字典值
	Description string `gorm:"column:description;size:255;not null;comment:描述;" json:"description"` // 描述
	Extra       string `gorm:"column:extra;size:255;not null;comment:额外信息;" json:"extra"`           // 额外信息
}

func (d DictDetail) TableName() string { return "sys_dict_detail" }
