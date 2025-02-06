package model

import (
	"lime/internal/common/model"
)

type CompileInfo struct {
	model.BaseModel
	CompileBase
}

type CompileBase struct {
	ProjectId     uint      `gorm:"column:project_id;size:200;not null;comment:项目ID;" json:"project_id"`   // 项目ID
	Output        string    `gorm:"column:output;size:-1;not null;comment:输出文件名;" json:"output"`           // 输出文件名
	Goos          OS_TYPE   `gorm:"column:goos;type:int;comment:目标操作系统;" json:"goos"`                      // 目标操作系统 0: windows 1: linux 2: darwin
	Goarch        ARCH_TYPE `gorm:"column:goarch;type:int;comment:目标架构;" json:"goarch"`                    // 目标架构 0: amd64 1: arm64
	Flags         model.Arr `gorm:"column:flags;size:-1;comment:编译标志;" json:"flags"`                       // 编译标志
	Ldflags       string    `gorm:"column:ldflags;size:-1;comment:链接标志;" json:"ldflags"`                   // 链接标志
	Tags          string    `gorm:"column:tags;size:-1;comment:目标架构;" json:"tags"`                         // 编译标签
	BeforeScripts Scripts   `gorm:"column:before_scripts;size:-1;comment:编译前执行的脚本;" json:"before_scripts"` // 编译前执行的脚本
	AfterScripts  Scripts   `gorm:"column:after_scripts;size:-1;comment:编译后执行的脚本;" json:"after_scripts"`   // 编译前执行的脚本
	EnvVars       model.KVS `gorm:"column:env_vars;size:-1;comment:编译期注入变量;" json:"env_vars"`              // 编译期注入变量 追加到链接标志中
}

func (CompileInfo) TableName() string {
	return "project_compile_info"
}
