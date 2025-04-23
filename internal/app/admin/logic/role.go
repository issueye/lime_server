package logic

import (
	"errors"
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	"lime/internal/app/admin/service"
	commonModel "lime/internal/common/model"
	"lime/internal/global"
	"strings"
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

func (lc *Role) GetNoHaveApis(condition *requests.RoleQryApi) ([]model.ApiInfo, error) {
	// 查询角色信息
	casbinSrv := service.NewCasbin(global.DB)
	casbinList, err := casbinSrv.GetRoleApis(condition.RoleCode)
	if err != nil {
		return nil, err
	}

	// 查询接口信息
	apiSrv := service.NewApiManage()
	apiList, err := apiSrv.ListFilter(condition)
	if err != nil {
		return nil, err
	}

	find := func(path string, method string) *requests.CasbinInfo {
		for _, v := range casbinList {
			objPath := strings.TrimPrefix(path, "/api/v1")
			if objPath == v.Path && v.Method == method {
				return &v
			}
		}

		return nil
	}

	var res []model.ApiInfo
	for _, v := range apiList {
		api := find(v.Path, v.Method.String())
		if api != nil {
			continue
		}

		res = append(res, *v)
	}

	return res, nil
}

func (lc *Role) GetRoleApis(condition *requests.RoleQryApi) ([]model.ApiInfo, error) {
	// 查询角色信息
	casbinSrv := service.NewCasbin(global.DB)
	casbinList, err := casbinSrv.GetRoleApis(condition.RoleCode)
	if err != nil {
		return nil, err
	}

	// 查询接口信息
	apiSrv := service.NewApiManage()
	apiList, err := apiSrv.ListFilter(condition)
	if err != nil {
		return nil, err
	}

	find := func(path string, method string) *model.ApiInfo {
		for _, v := range apiList {
			objPath := strings.TrimPrefix(v.Path, "/api/v1")
			if objPath == path && v.Method.String() == method {
				return v
			}
		}
		return nil
	}

	var res []model.ApiInfo
	for _, v := range casbinList {
		api := find(v.Path, v.Method)
		if api != nil {
			res = append(res, *api)
		}
	}

	return res, nil
}

func (lc *Role) RemoveApi(condition *requests.RoleApi) error {
	casbinSrv := service.NewCasbin(global.DB)
	path := strings.TrimPrefix(condition.Path, "/api/v1")
	return casbinSrv.RemoveRoleApi(condition.RoleCode, path, condition.Method)
}

func (lc *Role) AddApi(condition *requests.RoleApi) error {
	casbinSrv := service.NewCasbin(global.DB)
	path := strings.TrimPrefix(condition.Path, "/api/v1")
	return casbinSrv.AddRoleApi(condition.RoleCode, path, condition.Method)
}

func (lc *Role) FreshCasbin() error {
	casbinSrv := service.NewCasbin(global.DB)
	return casbinSrv.FreshCasbin()
}
