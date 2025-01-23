package logic

import (
	"lime/internal/app/project/model"
	"lime/internal/app/project/requests"
	"lime/internal/app/project/service"
	commonModel "lime/internal/common/model"
)

func CreateVersion(req *requests.CreateVersion) error {
	srv := service.NewVersion()
	return srv.Create(&req.VersionInfo)
}

func UpdateVersion(req *requests.UpdateVersion) error {
	return service.NewVersion().Update(req.ID, &req.VersionInfo)
}

func UpdateVersionBuildStatus(id uint, status model.BuildStatus) error {
	return service.NewVersion().UpdateVersionBuildStatus(id, status)
}

func DeleteVersion(id uint) error {
	return service.NewVersion().Delete(id)
}

func VersionList(condition *commonModel.PageQuery[*requests.QueryVersion]) (*commonModel.ResPage[model.VersionInfo], error) {
	return service.NewVersion().ListVersion(condition)
}
