package logic

import (
	"fmt"
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	"lime/internal/app/admin/service"
	commonModel "lime/internal/common/model"
)

func CreateDicts(req *requests.CreateDicts) error {
	srv := service.NewDicts()
	return srv.Create(&req.DictsInfo)
}

func UpdateDicts(req *requests.UpdateDicts) error {
	fmt.Println("UpdateDicts -> ", req)
	return service.NewDicts().Update(req.ID, &req.DictsInfo)
}

func DeleteDicts(id uint) error {
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

func DictsList(condition *commonModel.PageQuery[*requests.QueryDicts]) (*commonModel.ResPage[model.DictsInfo], error) {
	return service.NewDicts().ListDicts(condition)
}

func GetDicts(id uint) (*model.DictsInfo, error) {
	return service.NewDicts().GetById(id)
}

func GetDictsByCode(code string) (*model.DictsInfo, error) {
	return service.NewDicts().GetByField("code", code)
}

func SaveDetail(req *requests.SaveDetail) error {
	return service.NewDictDetail().Save(&req.DictDetail)
}

func DelDetail(id uint) error {
	return service.NewDictDetail().Delete(id)
}

func ListDetail(condition *commonModel.PageQuery[*requests.QueryDictsDetail]) (*commonModel.ResPage[model.DictDetail], error) {
	return service.NewDictDetail().List(condition)
}
