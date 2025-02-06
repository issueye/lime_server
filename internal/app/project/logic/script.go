package logic

import (
	"fmt"
	"lime/internal/app/project/model"

	"github.com/mattn/anko/vm"
)

func GetOutfileName(projectInfo model.ProjectInfo, versionInfo model.VersionInfo, info model.CompileInfo) (string, error) {
	if info.Output == "" {
		return "", fmt.Errorf("输出文件名称不能为空")
	}

	env := NewAnko()
	env.SetParams("compile", info)
	env.SetParams("project", projectInfo)
	env.SetParams("version", versionInfo)

	data, err := vm.Execute(env.GetEnv(), nil, info.Output)
	if err != nil {
		return "", err
	}

	rtnValue, ok := data.(string)
	if !ok {
		return "", fmt.Errorf("输出文件名称必须为字符串")
	}

	return rtnValue, nil
}
