package logic

import (
	"lime/internal/app/project/model"
	"lime/internal/app/project/requests"
	"lime/internal/app/project/service"
	commonModel "lime/internal/common/model"
)

func CreateProject(req *requests.CreateProject) error {
	srv := service.NewProject()
	return srv.Create(&req.ProjectInfo)
}

func UpdateProject(req *requests.UpdateProject) error {
	return service.NewProject().Update(req.ID, &req.ProjectInfo)
}

func DeleteProject(id uint) error {
	// 删除字典，并且删除对应明细数据
	srv := service.NewProject()

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

	branchSrv := service.NewBranch(srv.GetDB(), true)
	err = branchSrv.DeleteByFields(map[string]any{"project_id": info.ID})
	if err != nil {
		return err
	}

	tagSrv := service.NewTag(srv.GetDB(), true)
	err = tagSrv.DeleteByFields(map[string]any{"project_id": info.ID})
	if err != nil {
		return err
	}

	versionSrv := service.NewVersion(srv.GetDB(), true)
	err = versionSrv.DeleteByFields(map[string]any{"project_id": info.ID})
	if err != nil {
		return err
	}

	return nil
}

func ProjectList(condition *commonModel.PageQuery[*requests.QueryProject]) (*commonModel.ResPage[model.ProjectInfo], error) {
	return service.NewProject().ListProject(condition)
}

func GetProject(id uint) (*model.ProjectInfo, error) {
	return service.NewProject().GetById(id)
}
