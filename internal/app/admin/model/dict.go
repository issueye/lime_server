package model

import "lime/internal/common/model"

type DictsInfo struct {
	model.BaseModel
	DictsBase
}

type DictsBase struct {
	Code    string       `gorm:"column:code;size:200;not null;comment:菜单编码;" json:"code"`   // 菜单编码
	Name    string       `gorm:"column:name;size:200;not null;comment:菜单名称;" json:"name"`   // 菜单名称
	Remark  string       `gorm:"column:remark;size:255;not null;comment:备注;" json:"remark"` // 备注
	Details []DictDetail `gorm:"foreignKey:DictCode;references:code;" json:"details"`       // 字典详情
}

func (d DictsInfo) TableName() string { return "dicts_info" }

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

func (d DictDetail) TableName() string { return "dict_detail" }
