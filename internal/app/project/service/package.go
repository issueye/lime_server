package service

import (
	"lime/internal/app/project/model"
	"lime/internal/app/project/requests"
	commonModel "lime/internal/common/model"
	"lime/internal/common/service"

	"gorm.io/gorm"
)

type Package struct {
	service.BaseService[model.PackageInfo]
}

func NewPackage(args ...any) *Package {
	srv := &Package{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

func (r *Package) ListPackage(condition *commonModel.PageQuery[*requests.QueryPackage]) (*commonModel.ResPage[model.PackageInfo], error) {
	return service.GetList[model.PackageInfo](condition, func(qu *requests.QueryPackage, d *gorm.DB) *gorm.DB {
		if qu.ProjectId > 0 {
			d = d.Where("project_id = ?", qu.ProjectId)
		}
		if qu.VersionId > 0 {
			d = d.Where("version_id = ?", qu.VersionId)
		}
		if qu.Keywords != "" {
			d = d.Where("name like ?", "%"+qu.Keywords+"%")
		}
		return d
	})
}