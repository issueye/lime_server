package service

import (
	"lime/internal/app/project/model"
	"lime/internal/app/project/requests"
	commonModel "lime/internal/common/model"
	"lime/internal/common/service"

	"gorm.io/gorm"
)

type Project struct {
	service.BaseService[model.ProjectInfo]
}

func NewProject(args ...any) *Project {
	srv := &Project{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListProject
// 根据条件查询列表
func (r *Project) ListProject(condition *commonModel.PageQuery[*requests.QueryProject]) (*commonModel.ResPage[model.ProjectInfo], error) {
	return service.GetList[model.ProjectInfo](condition, func(qu *requests.QueryProject, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or remark like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		// 加载分支
		d = d.Preload("Branchs")
		// 加载版本
		d = d.Preload("Versions")
		// 加载标签
		d = d.Preload("Tags")

		return d
	})
}
