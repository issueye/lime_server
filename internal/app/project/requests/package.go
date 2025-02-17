package requests

import (
	"lime/internal/app/project/model"
	commonModel "lime/internal/common/model"
)

type CreatePackage struct {
	model.PackageInfo
}

func NewCreatePackage() *CreatePackage {
	return &CreatePackage{}
}

type QueryPackage struct {
	ProjectId uint   `json:"project_id"`
	VersionId uint   `json:"version_id"`
	Keywords  string `json:"keywords"`
}

func NewQueryPackage() *commonModel.PageQuery[*QueryPackage] {
	return commonModel.NewPageQuery[*QueryPackage](0, 0, &QueryPackage{})
}
