package model

import (
	"lime/internal/common/model"
	"time"
)

type PackageInfo struct {
	model.BaseModel
	PackageBase
}

type PackageBase struct {
	ProjectName  string    `gorm:"column:project_name;comment:项目名称" json:"project_name"`           // 项目名称
	ProjectId    uint      `gorm:"column:project_id;not null;comment:项目ID;" json:"project_id"`     // 项目ID
	Version      string    `gorm:"column:version;size:500;comment:版本" json:"version"`              // 版本
	VersionId    uint      `gorm:"column:version_id;not null;comment:版本ID;" json:"version_id"`     // 版本ID
	Name         string    `gorm:"column:name;size:200;not null;comment:文件名称;" json:"name"`        // 文件名称
	Path         string    `gorm:"column:path;size:500;not null;comment:文件路径;" json:"path"`        // 文件路径
	Size         int64     `gorm:"column:size;not null;comment:文件大小;" json:"size"`                 // 文件大小
	Hash         string    `gorm:"column:hash;size:64;not null;comment:文件hash;" json:"hash"`       // 文件hash
	DownloadNum  int       `gorm:"column:download_num;not null;comment:下载次数;" json:"download_num"` // 下载次数
	LastDownload time.Time `gorm:"column:last_download;comment:最后下载时间;" json:"last_download"`      // 最后下载时间
}

func (p PackageInfo) TableName() string { return "biz_project_package_info" }
