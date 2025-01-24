package model

import (
	"lime/internal/common/model"
	"time"
)

type ProjectInfo struct {
	model.BaseModel
	ProjectBase
}

type ProjectBase struct {
	Name         string        `gorm:"column:name;size:200;not null;comment:名称;" json:"name"`                         // 名称
	RepoUrl      string        `gorm:"column:repo_url;size:200;not null;comment:git 仓库地址;" json:"repo_url"`           // git 仓库地址
	RepoUser     string        `gorm:"column:repo_user;size:200;not null;comment:git 用户名称;" json:"repo_user"`         // git 用户名称
	RepoPassword string        `gorm:"column:repo_password;size:200;not null;comment:git 用户密码;" json:"repo_password"` // git 用户密码
	Description  string        `gorm:"column:description;size:200;not null;comment:项目描述;" json:"description"`         // 项目描述
	Branchs      []BranchInfo  `gorm:"foreignKey:ProjectId;references:id;" json:"branchs"`                            // 分支信息
	Tags         []TagInfo     `gorm:"foreignKey:ProjectId;references:id;" json:"tags"`                               // Tag 信息
	Versions     []VersionInfo `gorm:"foreignKey:ProjectId;references:id;" json:"versions"`                           // 版本信息
}

func (p ProjectInfo) TableName() string { return "biz_project_info" }

// 分支信息
type BranchInfo struct {
	model.BaseModel
	BranchBase
}

type BranchBase struct {
	ProjectId   uint   `gorm:"column:project_id;size:200;not null;comment:项目ID;" json:"project_id"`             // 项目ID
	Name        string `gorm:"column:name;size:200;not null;comment:分支名称;" json:"name"`                         // 分支名称
	Hash        string `gorm:"column:hash;size:200;not null;comment:Hash;" json:"hash"`                         // Hash 值
	Description string `gorm:"column:description;size:200;not null;comment:分支描述;" json:"description"`           // 分支描述
	Status      string `gorm:"column:status;size:50;not null;default:'developing';comment:分支状态;" json:"status"` // 分支状态: developing, merged, closed
}

func (p BranchInfo) TableName() string { return "biz_project_branch_info" }

// Tag 信息
type TagInfo struct {
	model.BaseModel
	TagBase
}

type TagBase struct {
	ProjectId     uint   `gorm:"column:project_id;size:200;not null;comment:项目ID;" json:"project_id"`                        // 项目ID
	Hash          string `gorm:"column:hash;size:200;not null;comment:Hash;" json:"hash"`                                    // Hash 值
	Name          string `gorm:"column:name;size:200;comment:标签名称;" json:"name"`                                             // 标签名称
	Description   string `gorm:"column:description;size:200;comment:标签描述;" json:"description"`                               // 标签描述
	ReleaseStatus string `gorm:"column:release_status;size:50;not null;default:'draft';comment:发布状态;" json:"release_status"` // 发布状态: draft, released
}

func (p TagInfo) TableName() string { return "biz_project_tag_info" }

type VersionInfo struct {
	model.BaseModel
	VersionBase
}

type BuildStatus string

const (
	BuildStatusPending  BuildStatus = "pending"
	BuildStatusBuilding BuildStatus = "building"
	BuildStatusSuccess  BuildStatus = "success"
	BuildStatusFailed   BuildStatus = "failed"
)

type VersionBase struct {
	ProjectId   uint        `gorm:"column:project_id;size:200;not null;comment:项目ID;" json:"project_id"`                      // 项目ID
	BranchName  string      `gorm:"column:branch_name;size:200;not null;comment:分支名称;" json:"branch_name"`                    // 分支名称
	BuildTime   time.Time   `gorm:"column:build_time;size:200;comment:构建时间;" json:"build_time"`                               // 构建时间
	Hash        string      `gorm:"column:hash;size:200;not null;comment:Hash;" json:"hash"`                                  // Hash 值
	Tag         string      `gorm:"column:tag;size:200;not null;comment:Tag;" json:"tag"`                                     // Tag 值
	Version     string      `gorm:"column:version;size:200;not null;comment:版本号;" json:"version"`                             // 版本号
	Description string      `gorm:"column:description;size:200;comment:版本描述;" json:"description"`                             // 版本描述
	BuildStatus BuildStatus `gorm:"column:build_status;size:50;not null;default:'pending';comment:构建状态;" json:"build_status"` // 构建状态
}

func (p VersionInfo) TableName() string { return "biz_project_version_info" }
