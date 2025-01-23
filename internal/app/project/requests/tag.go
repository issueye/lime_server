package requests

import (
	"encoding/json"
	"errors"
	"lime/internal/app/project/model"
	commonModel "lime/internal/common/model"
)

type CreateTag struct {
	model.TagInfo
}

func (req *CreateTag) Validate() error {
	if req.Name == "" {
		return errors.New("标签名称不能为空")
	}
	if req.ReleaseStatus != "draft" && req.ReleaseStatus != "released" {
		return errors.New("无效的发布状态")
	}
	return nil
}

func NewCreateTag() *CreateTag {
	return &CreateTag{
		TagInfo: model.TagInfo{},
	}
}

type UpdateTag struct {
	model.TagInfo
}

func (req *UpdateTag) Validate() error {
	if req.Name == "" {
		return errors.New("标签名称不能为空")
	}
	if req.ReleaseStatus != "draft" && req.ReleaseStatus != "released" {
		return errors.New("无效的发布状态")
	}
	return nil
}

func NewUpdateTag() *UpdateTag {
	return &UpdateTag{
		TagInfo: model.TagInfo{},
	}
}

func (req *CreateTag) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryTag struct {
	KeyWords string `json:"keywords" form:"keywords"` // 关键词
}

func NewQueryTag() *commonModel.PageQuery[*QueryTag] {
	return commonModel.NewPageQuery(0, 0, &QueryTag{})
}
