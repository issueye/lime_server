package logic

import (
	"fmt"
	"lime/internal/app/project/model"
)

// InvokeParams 执行参数
type InvokeParams struct {
	WorkDir string            `json:"workDir"` // 工作目录
	Code    string            `json:"code"`    // 代码
	Project model.ProjectInfo `json:"project"` // 项目信息
	Version model.VersionInfo `json:"version"` // 版本信息
	Compile model.CompileInfo `json:"compile"` // 编译信息
}

func GetOutfileName(params InvokeParams) (string, error) {
	data, err := runCode(params)
	if err != nil {
		return "", err
	}

	rtnValue, ok := data.(string)
	if !ok {
		return "", fmt.Errorf("输出文件名称必须为字符串")
	}

	return rtnValue, nil
}

func BeforeScript(params InvokeParams) error {
	_, err := runCode(params)
	if err != nil {
		return err
	}

	return nil
}

func AfterScript(params InvokeParams) error {
	_, err := runCode(params)
	if err != nil {
		return err
	}

	return nil
}

func InjectEnv(params InvokeParams) (string, error) {
	data, err := runCode(params)
	if err != nil {
		return "", err
	}

	rtnValue, ok := data.(string)
	if !ok {
		return "", fmt.Errorf("输出内容必须为字符串")
	}

	return rtnValue, nil
}

func runCode(params InvokeParams) (any, error) {
	if params.Code == "" {
		return nil, fmt.Errorf("代码不能为空")
	}

	env := NewAnko()
	env.SetParams("workDir", params.WorkDir)
	env.SetParams("compile", params.Compile)
	env.SetParams("project", params.Project)
	env.SetParams("version", params.Version)
	data, err := env.Execute(params.Code)
	if err != nil {
		return nil, err
	}

	return data, nil
}
