package logic

import (
	"fmt"
	"lime/internal/app/project/model"
	"lime/internal/app/project/requests"
	"lime/internal/app/project/service"
	commonModel "lime/internal/common/model"
	"os"
	"time"
)

func CreatePackage(req *requests.CreatePackage) error {
	return service.NewPackage().Create(&req.PackageInfo)
}

func DeletePackage(id uint) error {
	srv := service.NewPackage()

	info, err := srv.GetById(id)
	if err != nil {
		return err
	}

	// 删除物理文件
	if err := os.Remove(info.Path); err != nil {
		return fmt.Errorf("删除文件失败: %v", err)
	}

	return srv.Delete(id)
}

func PackageList(condition *commonModel.PageQuery[*requests.QueryPackage]) (*commonModel.ResPage[model.PackageInfo], error) {
	return service.NewPackage().ListPackage(condition)
}

func DownloadPackage(id uint) (string, error) {
	srv := service.NewPackage()

	info, err := srv.GetById(id)
	if err != nil {
		return "", err
	}

	// 更新下载信息
	info.DownloadNum++
	info.LastDownload = time.Now()

	if err := srv.Update(id, info); err != nil {
		return "", err
	}

	return info.Path, nil
}
