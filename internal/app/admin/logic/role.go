package logic

import (
	"errors"
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	"lime/internal/app/admin/service"
	commonModel "lime/internal/common/model"
	commonSrv "lime/internal/common/service"
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

	srv := service.NewRole()
	srv.Begin()

	var err error
	defer func() {
		if err != nil {
			srv.Rollback()
			return
		}

		srv.Commit()
	}()

	err = srv.UpdateByMap(uint(r.Id), data)
	if err != nil {
		return err
	}

	roleMenuSrv := service.NewRoleMenu(srv.GetDB())
	// 删除角色的所有菜单
	err = roleMenuSrv.DeleteByFields(map[string]any{"role_code": r.Code})
	if err != nil {
		return err
	}

	// 新增角色的菜单
	datas := make([]model.RoleMenu, 0)
	for _, v := range r.MenuCodes {
		datas = append(datas, model.RoleMenu{
			RoleCode: r.Code,
			MenuCode: v,
		})
	}

	err = roleMenuSrv.CreateBatch(datas)
	if err != nil {
		return err
	}

	return nil
}

// 根据ID查询数据
func (lc *Role) GetRoleById(id uint) (*model.Role, error) {
	return service.NewRole().GetById(id)
}

// 根据条件查询数据
func (lc *Role) ListRole(condition *commonModel.PageQuery[*requests.QueryRole]) (*commonModel.ResPage[model.Role], error) {
	return service.NewRole().ListRole(condition)
}

func (lc *Role) List(condition *requests.QueryRole) ([]*model.Role, error) {
	return service.NewRole().List(condition)
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

func (lc *Role) AddApi(condition *requests.CreateRoleApis) error {
	casbinSrv := service.NewCasbin(global.DB)

	// 查询所有的接口信息
	apiSrv := service.NewApiManage()
	conditions := []commonSrv.Condition{
		{Field: "id", Value: condition.Apis, Exp: "in"},
	}
	apis, err := apiSrv.GetDatasByFields(conditions)
	if err != nil {
		return err
	}

	datas := make([]service.RoleApiData, 0)
	for _, v := range apis {
		path := strings.TrimPrefix(v.Path, "/api/v1")
		datas = append(datas, service.RoleApiData{
			RoleCode: condition.RoleCode,
			Path:     path,
			Method:   v.Method.String(),
		})
	}

	err = casbinSrv.AddRoleApis(datas)
	if err != nil {
		return err
	}

	return nil
}

func (lc *Role) FreshCasbin() error {
	casbinSrv := service.NewCasbin(global.DB)
	return casbinSrv.FreshCasbin()
}
