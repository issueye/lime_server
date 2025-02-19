package model

import "lime/internal/common/model"

type DictsInfo struct {
	model.BaseModel
	DictsBase
}

type ContentType int

const (
	ContentTypeJson       ContentType = 1 // json
	ContentTypeText       ContentType = 2 // text
	ContentTypeAnkoScript ContentType = 3 // anko脚本
)

type DictsBase struct {
	Code        string       `gorm:"column:code;size:200;not null;comment:编码;" json:"code"`          // 编码
	Name        string       `gorm:"column:name;size:200;not null;comment:名称;" json:"name"`          // 名称
	ContentType ContentType  `gorm:"column:content_type;not null;comment:内容类型;" json:"content_type"` // 内容类型
	Remark      string       `gorm:"column:remark;size:255;not null;comment:备注;" json:"remark"`      // 备注
	Details     []DictDetail `gorm:"foreignKey:DictCode;references:code;" json:"details"`            // 字典详情
}

func (d DictsInfo) TableName() string { return "sys_dict_info" }

type DictDetail struct {
	model.BaseModel
	DictDetailBase
}

type DictDetailBase struct {
	DictCode string `gorm:"column:dict_code;size:200;not null;comment:字典编码;" json:"dict_code"` // 字典编码
	Key      string `gorm:"column:key;size:200;not null;comment:字典键;" json:"key"`              // 字典键
	Value    string `gorm:"column:value;size:200;not null;comment:字典值;" json:"val"`            // 字典值
	Remark   string `gorm:"column:remark;size:255;not null;comment:备注;" json:"remark"`         // 备注
}

func (d DictDetail) TableName() string { return "sys_dict_detail" }
