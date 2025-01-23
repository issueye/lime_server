package requests

import (
	"encoding/json"
	"errors"
	"lime/internal/app/project/model"
	commonModel "lime/internal/common/model"
)

type CreateBranch struct {
	model.BranchInfo
}

func (req *CreateBranch) Validate() error {
	if req.Name == "" {
		return errors.New("分支名称不能为空")
	}
	if req.Status != "developing" && req.Status != "merged" && req.Status != "closed" {
		return errors.New("无效的分支状态")
	}
	return nil
}

func NewCreateBranch() *CreateBranch {
	return &CreateBranch{
		BranchInfo: model.BranchInfo{},
	}
}

type UpdateBranch struct {
	model.BranchInfo
}

func (req *UpdateBranch) Validate() error {
	if req.Name == "" {
		return errors.New("分支名称不能为空")
	}
	if req.Status != "developing" && req.Status != "merged" && req.Status != "closed" {
		return errors.New("无效的分支状态")
	}
	return nil
}

func NewUpdateBranch() *UpdateBranch {
	return &UpdateBranch{
		BranchInfo: model.BranchInfo{},
	}
}

func (req *CreateBranch) ToJson() string {
	data, err := json.Marshal(req)
	if err != nil {
		return ""
	}
	return string(data)
}

type QueryBranch struct {
	KeyWords string `json:"keywords" form:"keywords"` // 关键词
}

func NewQueryBranch() *commonModel.PageQuery[*QueryBranch] {
	return commonModel.NewPageQuery(0, 0, &QueryBranch{})
}
