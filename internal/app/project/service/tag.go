package service

import (
	"errors"
	"lime/internal/app/project/model"
	"lime/internal/app/project/requests"
	commonModel "lime/internal/common/model"
	"lime/internal/common/service"

	"gorm.io/gorm"
)

type Tag struct {
	service.BaseService[model.TagInfo]
}

func NewTag(args ...any) *Tag {
	srv := &Tag{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListTag
// 根据条件查询列表
func (r *Tag) ListTag(condition *commonModel.PageQuery[*requests.QueryTag]) (*commonModel.ResPage[model.TagInfo], error) {
	return service.GetList[model.TagInfo](condition, func(qu *requests.QueryTag, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or remark like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}

// UpdateTagReleaseStatus 更新标签发布状态
func (r *Tag) UpdateTagReleaseStatus(id uint, status string) error {
	if status != "draft" && status != "released" {
		return errors.New("无效的发布状态")
	}

	return r.DB.Model(&model.TagInfo{}).
		Where("id = ?", id).
		Update("release_status", status).Error
}
