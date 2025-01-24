package model

import "lime/internal/common/model"

type CompileInfo struct {
	model.BaseModel
	CompileBase
}

type Script struct {
	Name    string `json:"name"`    // 脚本名称
	Content string `json:"content"` // 脚本内容
}

type EnvVar struct {
	Key   string `json:"key"`   // 环境变量键
	Value string `json:"value"` // 环境变量值
}

type CompileBase struct {
	Output  string   `json:"output"`   // 输出文件名
	Goos    string   `json:"goos"`     // 目标操作系统
	Goarch  string   `json:"goarch"`   // 目标架构
	Flags   []string `json:"flags"`    // 编译标志
	Ldflags string   `json:"ldflags"`  // 链接标志
	Tags    string   `json:"tags"`     // 编译标签
	Scripts []Script `json:"scripts"`  // 编译前执行的JavaScript脚本
	EnvVars []EnvVar `json:"env_vars"` // 环境变量
}
