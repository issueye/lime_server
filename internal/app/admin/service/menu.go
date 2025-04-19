package service

import (
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	commonModel "lime/internal/common/model"
	"lime/internal/common/service"
	"lime/internal/global"

	"gorm.io/gorm"
)

type Menu struct {
	service.BaseService[model.Menu]
}

func NewMenu(args ...any) *Menu {
	srv := &Menu{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// AddMenus 添加菜单
func (m *Menu) AddMenu(menu *model.Menu) error {
	return global.DB.Create(menu).Error
}

// 检查菜单是否存在
func (m *Menu) CheckMenuExist(menu *model.Menu) (bool, error) {
	var count int64
	err := global.DB.Model(&model.Menu{}).Where("code = ?", menu.Code).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ListMenu
// 根据条件查询列表
func (m *Menu) ListMenu(condition *commonModel.PageQuery[*requests.QueryMenu]) (*commonModel.ResPage[model.Menu], error) {
	return service.GetList[model.Menu](condition, func(qu *requests.QueryMenu, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or description like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		return d
	})
}

func (m *Menu) GetMenuByCodes(codes []string) ([]*model.Menu, error) {
	list := make([]*model.Menu, 0)

	qry := m.GetDB().Where("code in (?)", codes)
	err := qry.Find(&list).Error
	return list, err
}
