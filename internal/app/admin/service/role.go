package service

import (
	"fmt"
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	commonModel "lime/internal/common/model"
	"lime/internal/common/service"
	"lime/internal/global"

	"gorm.io/gorm"
)

type Role struct {
	service.BaseService[model.Role]
}

func NewRole(args ...any) *Role {
	srv := &Role{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}

// ListRole
// 根据条件查询列表
func (r *Role) ListRole(condition *commonModel.PageQuery[*requests.QueryRole]) (*commonModel.ResPage[model.Role], error) {
	return service.GetList[model.Role](condition, func(qu *requests.QueryRole, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or remark like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		d.Preload("RoleMenus")

		return d
	})
}

func (r *Role) List(condition *requests.QueryRole) ([]*model.Role, error) {
	return service.GetDatas[model.Role](condition, func(qu *requests.QueryRole, d *gorm.DB) *gorm.DB {
		if qu.KeyWords != "" {
			d = d.Where("name like ? or remark like ?", "%"+qu.KeyWords+"%", "%"+qu.KeyWords+"%")
		}

		d.Preload("RoleMenus")

		return d
	})
}

func (r *Role) GetRoleMenus(Role_code string) ([]*model.Menu, error) {
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

func (r *Role) SaveRoleMenus(Role_code string, menu_codes []string) error {
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
