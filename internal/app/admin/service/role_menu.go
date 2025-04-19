package service

import (
	"fmt"
	"lime/internal/app/admin/model"
	"lime/internal/common/service"
	"lime/internal/global"

	"gorm.io/gorm"
)

type RoleMenu struct {
	service.BaseService[model.RoleMenu]
}

func NewRoleMenu(args ...any) *RoleMenu {
	srv := &RoleMenu{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

func (r *RoleMenu) GetRoleMenus(Role_code string) ([]*model.Menu, error) {
	menu := make([]*model.Menu, 0)

	rm := global.DB.Model(&model.RoleMenu{})
	if Role_code != "" {
		rm = rm.Where("role_code =?", Role_code)
	}

	sqlStr := rm.ToSQL(func(tx *gorm.DB) *gorm.DB { return tx.Find(nil) })
	qry := global.DB.Model(&model.Menu{}).Joins(fmt.Sprintf(`left join (%s) rm on rm.menu_code = sys_menu.code`, sqlStr)).
		Select("sys_menu.*,case when rm.role_code is not null then 1 else 0 end as is_have")

	err := qry.Find(&menu).Error
	return menu, err
}

func (r *RoleMenu) SaveRoleMenus(Role_code string, menu_codes []string) error {
	rm := global.DB.Model(&model.RoleMenu{}).Where("role_code =?", Role_code)
	err := rm.Delete(&model.RoleMenu{}).Error
	if err != nil {
		return err
	}

	rm = global.DB.Model(&model.RoleMenu{})
	for _, code := range menu_codes {
		rm = rm.Create(&model.RoleMenu{
			RoleCode: Role_code,
			MenuCode: code,
		})
	}

	return rm.Error
}
