package service

import (
	"errors"
	"lime/internal/app/project/model"
	"lime/internal/app/project/requests"
	commonModel "lime/internal/common/model"
	"lime/internal/common/service"

	"gorm.io/gorm"
)

type Version struct {
	service.BaseService[model.VersionInfo]
}

func NewVersion(args ...any) *Version {
	srv := &Version{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListVersion
// 根据条件查询列表
func (r *Version) ListVersion(condition *commonModel.PageQuery[*requests.QueryVersion]) (*commonModel.ResPage[model.VersionInfo], error) {
	return service.GetList[model.VersionInfo](condition, func(qu *requests.QueryVersion, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("branch_name like ? or description like ? or version like ?",
				"%"+qu.KeyWords+"%",
				"%"+qu.KeyWords+"%",
				"%"+qu.KeyWords+"%",
			)
		}

		if qu.ProjectId != 0 {
			d = d.Where("project_id =?", qu.ProjectId)
		}

		return d
	})
}

// UpdateVersionBuildStatus 更新版本构建状态
func (r *Version) UpdateVersionBuildStatus(id uint, status model.BuildStatus) error {
	switch status {
	case model.BuildStatusPending, model.BuildStatusBuilding,
		model.BuildStatusSuccess, model.BuildStatusFailed:
		return r.DB.Model(&model.VersionInfo{}).
			Where("id = ?", id).
			Update("build_status", status).Error
	default:
		return errors.New("无效的构建状态")
	}
}
