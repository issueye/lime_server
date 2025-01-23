package service

import (
	"errors"
	"lime/internal/app/project/model"
	"lime/internal/app/project/requests"
	commonModel "lime/internal/common/model"
	"lime/internal/common/service"

	"gorm.io/gorm"
)

type Branch struct {
	service.BaseService[model.BranchInfo]
}

func NewBranch(args ...any) *Branch {
	srv := &Branch{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListBranch
// 根据条件查询列表
func (r *Branch) ListBranch(condition *commonModel.PageQuery[*requests.QueryBranch]) (*commonModel.ResPage[model.BranchInfo], error) {
	return service.GetList[model.BranchInfo](condition, func(qu *requests.QueryBranch, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or remark like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}

// UpdateBranchStatus 更新分支状态
func (r *Branch) UpdateBranchStatus(id uint, status string) error {
	if status != "developing" && status != "merged" && status != "closed" {
		return errors.New("无效的分支状态")
	}

	return r.DB.Model(&model.BranchInfo{}).
		Where("id = ?", id).
		Update("status", status).Error
}
