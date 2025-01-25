package requests

import (
	"encoding/json"
	"errors"
	"lime/internal/app/project/model"
	commonModel "lime/internal/common/model"
)

type CreateVersion struct {
	model.VersionInfo
}

func (req *CreateVersion) Validate() error {
	if req.Version == "" {
		return errors.New("版本号不能为空")
	}
	if req.BuildStatus != "pending" && req.BuildStatus != "building" &&
		req.BuildStatus != "success" && req.BuildStatus != "failed" {
		return errors.New("无效的构建状态")
	}
	return nil
}

func NewCreateVersion() *CreateVersion {
	return &CreateVersion{
		VersionInfo: model.VersionInfo{},
	}
}

type UpdateVersion struct {
	model.VersionInfo
}

func (req *UpdateVersion) Validate() error {
	if req.Version == "" {
		return errors.New("版本号不能为空")
	}
	if req.BuildStatus != "pending" && req.BuildStatus != "building" &&
		req.BuildStatus != "success" && req.BuildStatus != "failed" {
		return errors.New("无效的构建状态")
	}
	return nil
}

func NewUpdateVersion() *UpdateVersion {
	return &UpdateVersion{
		VersionInfo: model.VersionInfo{},
	}
}

func (req *CreateVersion) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryVersion struct {
	KeyWords  string `json:"keywords" form:"keywords"`     // 关键词
	ProjectId uint   `json:"project_id" form:"project_id"` // 项目ID
}

func NewQueryVersion() *commonModel.PageQuery[*QueryVersion] {
	return commonModel.NewPageQuery(0, 0, &QueryVersion{})
}
