package requests

import (
	"encoding/json"
	"lime/internal/app/project/model"
	commonModel "lime/internal/common/model"
)

type CreateProject struct {
	model.ProjectInfo
}

func NewCreateProject() *CreateProject {
	return &CreateProject{
		ProjectInfo: model.ProjectInfo{},
	}
}

type UpdateProject struct {
	model.ProjectInfo
}

func NewUpdateProject() *UpdateProject {
	return &UpdateProject{
		ProjectInfo: model.ProjectInfo{},
	}
}

func (req *CreateProject) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryProject struct {
	KeyWords string `json:"keywords" form:"keywords"` // 关键词
}

func NewQueryProject() *commonModel.PageQuery[*QueryProject] {
	return commonModel.NewPageQuery(0, 0, &QueryProject{})
}
