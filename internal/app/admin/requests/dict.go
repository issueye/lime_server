package requests

import (
	"encoding/json"
	"lime/internal/app/admin/model"
	commonModel "lime/internal/common/model"
)

type CreateDicts struct {
	model.DictsInfo
}

func NewCreateDicts() *CreateDicts {
	return &CreateDicts{
		DictsInfo: model.DictsInfo{},
	}
}

type UpdateDicts struct {
	model.DictsInfo
}

func NewUpdateDicts() *UpdateDicts {
	return &UpdateDicts{
		DictsInfo: model.DictsInfo{},
	}
}

func (req *CreateDicts) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryDicts struct {
	KeyWords string `json:"keywords" form:"keywords"` // 关键词
}

func NewQueryDicts() *commonModel.PageQuery[*QueryDicts] {
	return commonModel.NewPageQuery(0, 0, &QueryDicts{})
}

type SaveDetail struct {
	model.DictDetail
}

func NewSaveDetail() *SaveDetail {
	return &SaveDetail{
		DictDetail: model.DictDetail{},
	}
}

type QueryDictsDetail struct {
	KeyWords string `json:"keywords" form:"keywords"`   // 关键词
	DictCode string `json:"dict_code" form:"dict_code"` // 字典编码
}

func NewQueryDictsDetail() *commonModel.PageQuery[*QueryDictsDetail] {
	return commonModel.NewPageQuery(0, 0, &QueryDictsDetail{})
}
