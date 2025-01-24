package logic

import (
	"lime/internal/app/project/requests"
	"lime/internal/app/project/service"
)

// CompileProject 编译项目
func CompileProject(req *requests.CompileRequest) error {
	info := service.CompileInfo{
		Output:  req.Output,
		Goos:    req.Goos,
		Goarch:  req.Goarch,
		Flags:   req.Flags,
		Ldflags: req.Ldflags,
		Tags:    req.Tags,
		Scripts: req.Scripts,
		EnvVars: req.EnvVars,
	}

	// 保存编译信息
	if err := service.SaveCompileInfo(info); err != nil {
		return err
	}

	// 执行编译
	return service.CompileProgram(info)
}
