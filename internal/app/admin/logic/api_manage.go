package logic

import (
	"fmt"
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	"lime/internal/app/admin/service"
	commonModel "lime/internal/common/model"
	"lime/internal/global"
)

type ApiManageLogic struct{}

func NewApiManageLogic() *ApiManageLogic {
	return &ApiManageLogic{}
}

func (lc *ApiManageLogic) IsNotExistAdd(menu *model.ApiInfo) {
	srv := service.NewApiManage()

	isHave, err := srv.CheckExist(menu)
	if err != nil {
		global.Logger.Sugar().Errorf("检查接口信息是否存在失败: %s", err.Error())
		return
	}

	if !isHave {
		err = srv.Create(menu)
		if err != nil {
			global.Logger.Sugar().Errorf("添加接口信息失败: %s", err.Error())
		}
	}
}

func (lc *ApiManageLogic) Create(info *requests.CreateApiInfo) error {
	srv := service.NewApiManage()

	// 检查是否存在
	apiDataInfo, err := srv.GetByMap(map[string]any{
		"path":   info.Path,
		"method": info.Method,
		"group":  info.Group,
	})
	if err != nil {
		return err
	}

	if apiDataInfo.ID > 0 {
		return fmt.Errorf("【%s-%s】接口信息已存在", info.Path, info.Method)
	}

	base := model.ApiBase{
		Title:  info.Title,
		Path:   info.Path,
		Method: info.Method,
		Group:  info.Group,
	}

	apiInfo := model.NewApi(base)
	return srv.Create(apiInfo)
}

func (lc *ApiManageLogic) Update(info *requests.UpdateApiInfo) error {
	srv := service.NewApiManage()

	// 检查是否存在
	apiDataInfo, err := srv.GetByMap(map[string]any{
		"path":   info.Path,
		"method": info.Method,
		"group":  info.Group,
	})
	if err != nil {
		return err
	}

	if apiDataInfo.ID == 0 {
		return fmt.Errorf("【%s-%s】接口信息不存在", info.Path, info.Method)
	}

	updateData := map[string]any{
		"title":  info.Title,
		"path":   info.Path,
		"group":  info.Group,
		"method": info.Method,
	}

	return srv.UpdateByMap(info.Id, updateData)
}

func (lc *ApiManageLogic) Delete(id uint) error {
	srv := service.NewApiManage()
	return srv.Delete(id)
}

func (lc *ApiManageLogic) GetList(condition *commonModel.PageQuery[*requests.QueryApiInfo]) (*commonModel.ResPage[model.ApiInfo], error) {
	return service.NewApiManage().GetList(condition)
}
