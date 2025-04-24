package logic

import (
	"fmt"
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	"lime/internal/app/admin/service"
	commonModel "lime/internal/common/model"
)

type DictsLogic struct{}

func NewDictsLogic() *DictsLogic {
	return &DictsLogic{}
}

func (lc *DictsLogic) CreateDicts(req *requests.CreateDicts) error {
	srv := service.NewDicts()
	return srv.Create(&req.DictsInfo)
}

func (lc *DictsLogic) UpdateDicts(req *requests.UpdateDicts) error {
	fmt.Println("UpdateDicts -> ", req)
	return service.NewDicts().Update(req.ID, &req.DictsInfo)
}

func (lc *DictsLogic) DeleteDicts(id uint) error {
	// 删除字典，并且删除对应明细数据
	srv := service.NewDicts()

	info, err := srv.GetById(id)
	if err != nil {
		return err
	}

	srv.Begin()

	defer func() {
		if err != nil {
			srv.Rollback()
			return
		}

		srv.Commit()
	}()

	err = srv.Delete(id)
	if err != nil {
		return err
	}

	detailSrv := service.NewDictDetail(srv.GetDB(), true)
	err = detailSrv.DeleteByFields(map[string]any{"dict_code": info.Code})
	return err
}

func (lc *DictsLogic) DictsList(condition *commonModel.PageQuery[*requests.QueryDicts]) (*commonModel.ResPage[model.DictsInfo], error) {
	return service.NewDicts().ListDicts(condition)
}

func (lc *DictsLogic) GetDicts(id uint) (*model.DictsInfo, error) {
	return service.NewDicts().GetById(id)
}

func (lc *DictsLogic) GetDictsByCode(code string) (*model.DictsInfo, error) {
	return service.NewDicts().GetByField("code", code)
}

func (lc *DictsLogic) SaveDetail(req *requests.SaveDetail) error {
	return service.NewDictDetail().Save(&req.DictDetail)
}

func (lc *DictsLogic) DelDetail(key string) error {
	return service.NewDictDetail().DeleteByFields(map[string]any{"key": key})
}

func (lc *DictsLogic) ListDetail(condition *commonModel.PageQuery[*requests.QueryDictsDetail]) (*commonModel.ResPage[model.DictDetail], error) {
	return service.NewDictDetail().List(condition)
}

func (lc *DictsLogic) IsNotExistAdd(req model.DictsInfo) error {
	dictsSrv := service.NewDicts()

	count, err := dictsSrv.GetCountByFields(map[string]any{"code": req.Code})
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("字典编码[%s]已存在", req.Code)
	}

	// 创建字典
	err = dictsSrv.Create(&req)
	if err != nil {
		return err
	}

	return nil
}
