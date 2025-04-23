package logic

import (
	"errors"
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	"lime/internal/app/admin/service"
	commonModel "lime/internal/common/model"
)

type Role struct{}

func NewRoleLogic() *Role {
	return &Role{}
}

// 创建数据
func (lc *Role) CreateRole(r *requests.CreateRole) error {

	srv := service.NewRole()

	data, err := srv.GetByField("code", r.Code)
	if err != nil {
		return err
	}

	if data.ID != 0 {
		return errors.New("角色编码已存在")
	}

	info := &model.Role{
		RoleBase: model.RoleBase{
			Code:   r.Code,
			Name:   r.Name,
			Remark: r.Remark,
		},
	}

	return service.NewRole().Create(info)
}

// 更新数据
func (lc *Role) UpdateRole(r *requests.UpdateRole) error {
	data := make(map[string]any)
	data["code"] = r.Code
	data["name"] = r.Name
	data["remark"] = r.Remark

	return service.NewRole().UpdateByMap(uint(r.Id), data)
}

// 根据ID查询数据
func (lc *Role) GetRoleById(id uint) (*model.Role, error) {
	return service.NewRole().GetById(id)
}

// 根据条件查询数据
func (lc *Role) ListRole(condition *commonModel.PageQuery[*requests.QueryRole]) (*commonModel.ResPage[model.Role], error) {
	return service.NewRole().ListRole(condition)
}

// 删除数据
func (lc *Role) DeleteRole(id uint) error {
	return service.NewRole().Delete(id)
}
