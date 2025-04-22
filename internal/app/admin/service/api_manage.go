package service

import (
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	commonModel "lime/internal/common/model"
	"lime/internal/common/service"
	"lime/internal/global"

	"gorm.io/gorm"
)

type ApiManage struct {
	service.BaseService[model.ApiInfo]
}

func NewApiManage(args ...any) *ApiManage {
	srv := &ApiManage{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// CheckExist
// 检查菜单是否存在
func (srv *ApiManage) CheckExist(api *model.ApiInfo) (bool, error) {
	var count int64
	err := global.DB.Model(&model.ApiInfo{}).Where("path = ? and method = ?", api.Path, api.Method).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ListApiInfo
// 根据条件查询列表
func (srv *ApiManage) GetList(condition *commonModel.PageQuery[*requests.QueryApiInfo]) (*commonModel.ResPage[model.ApiInfo], error) {
	return service.GetList[model.ApiInfo](condition, func(qu *requests.QueryApiInfo, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("title like ? or path like ? or method like ?",
				"%"+qu.KeyWords+"%",
				"%"+qu.KeyWords+"%",
				"%"+qu.KeyWords+"%",
			)
		}

		return d
	})
}
