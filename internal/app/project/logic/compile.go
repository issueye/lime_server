package logic

import (
	"fmt"
	"lime/internal/app/project/model"
	"lime/internal/app/project/repo"
	"lime/internal/app/project/requests"
	"lime/internal/app/project/service"
	"lime/internal/app/websocket"
	"os"
	"os/exec"
)

func SaveCompileConfig(req *requests.SaveCompileConfigRequest) error {
	srv := service.NewCompile()
	err := srv.DeleteByFields(map[string]any{"project_id": req.ProjectId})
	if err != nil {
		return err
	}

	return srv.Create(&req.CompileInfo)
}

func GetConfigByProjectId(projectId uint) (*model.CompileInfo, error) {
	srv := service.NewCompile()
	info, err := srv.GetByField("project_id", projectId)
	if err != nil {
		return nil, err
	}

	return info, nil
}

// CompileProject 编译项目
func CompileProject(projectId uint, versionId uint) error {
	req, err := GetConfigByProjectId(projectId)
	if err != nil {
		return err
	}

	// 检查编译配置是否存在
	if req == nil {
		return fmt.Errorf("编译配置不存在")
	}

	// 获取项目信息
	projectSrv := service.NewProject()
	projectInfo, err := projectSrv.GetByMap(map[string]any{"id": projectId})
	if err != nil {
		return err
	}

	// 获取版本信息
	versionSrv := service.NewVersion()
	versionInfo, err := versionSrv.GetByMap(map[string]any{"id": versionId, "project_id": projectId})
	if err != nil {
		return err
	}

	// 执行编译
	return CompileProgram(*projectInfo, *versionInfo, *req)
}

func CompileProgram(projectInfo model.ProjectInfo, versionInfo model.VersionInfo, info model.CompileInfo) error {
	// 获取输出文件名称
	SendMessage(versionInfo, "获取输出文件名称")
	output, err := GetOutfileName(
		InvokeParams{
			Code:    info.Output,
			Project: projectInfo,
			Version: versionInfo,
			Compile: info,
		},
	)
	if err != nil {
		SendMessage(versionInfo, fmt.Sprintf("获取输出文件名称失败: %v", err))
		return fmt.Errorf("获取输出文件名称失败: %v", err)
	}

	SendMessage(versionInfo, fmt.Sprintf("输出文件名称: %s", output))

	SendMessage(versionInfo, "拉取代码到临时目录")
	// 拉取代码到临时目录
	gitRepo, err := repo.PullCode(&projectInfo, versionInfo)
	if err != nil {
		SendMessage(versionInfo, fmt.Sprintf("拉取代码失败: %v", err))
		return fmt.Errorf("拉取代码失败: %v", err)
	}

	// checkout 到指定版本
	SendMessage(versionInfo, fmt.Sprintf("切换代码版本: %s  -hash:%s", versionInfo.Version, versionInfo.Hash))
	err = repo.Checkout(versionInfo, gitRepo)
	if err != nil {
		SendMessage(versionInfo, fmt.Sprintf("切换代码版本失败: %v", err))
		return fmt.Errorf("切换代码版本失败: %v", err)
	}

	// 获取仓库工作目录
	workDir, err := gitRepo.Worktree()
	if err != nil {
		SendMessage(versionInfo, fmt.Sprintf("获取工作目录失败: %v", err))
		return fmt.Errorf("获取工作目录失败: %v", err)
	}

	// 执行编译前的脚本
	SendMessage(versionInfo, "执行编译前的脚本")
	if err := executeScripts(
		workDir.Filesystem.Root(),
		info.BeforeScripts,
		projectInfo,
		versionInfo,
		info,
	); err != nil {
		SendMessage(versionInfo, fmt.Sprintf("执行脚本失败: %v", err))
	}

	// 执行构建命令
	SendMessage(versionInfo, "执行构建命令")
	err = runCommand(output, workDir.Filesystem.Root(), info, projectInfo, versionInfo)
	if err != nil {
		SendMessage(versionInfo, fmt.Sprintf("构建命令失败: %v", err))
		return fmt.Errorf("构建命令失败: %v", err)
	}

	// 执行编译后的脚本
	SendMessage(versionInfo, "执行编译后的脚本")
	if err := executeScripts(
		workDir.Filesystem.Root(),
		info.AfterScripts,
		projectInfo,
		versionInfo,
		info,
	); err != nil {
		SendMessage(versionInfo, fmt.Sprintf("执行脚本失败: %v", err))
		return fmt.Errorf("执行脚本失败: %v", err)
	}

	SendMessage(versionInfo, "编译完成")

	return nil
}

// executeScripts 执行编译前的JavaScript脚本
func executeScripts(workDir string, scripts model.Scripts, projectInfo model.ProjectInfo, versionInfo model.VersionInfo, info model.CompileInfo) error {
	for _, script := range scripts {
		fmt.Println("脚本内容：", script.Content)
		err := BeforeScript(
			InvokeParams{
				WorkDir: workDir,
				Code:    script.Content,
				Project: projectInfo,
				Version: versionInfo,
				Compile: info,
			},
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func runCommand(output string, workDir string, info model.CompileInfo, projectInfo model.ProjectInfo, versionInfo model.VersionInfo) error {
	// 准备编译命令
	args := []string{"build", "-o", output}

	// 设置编译标志
	ldflags := info.Ldflags
	if len(info.EnvVars) > 0 {
		for _, env := range info.EnvVars {
			if env.Key == "" || env.Value == "" {
				continue
			}

			// 获取环境变量的值
			val, _ := InjectEnv(
				InvokeParams{
					Code:    env.Value,
					Project: projectInfo,
					Version: versionInfo,
					Compile: info,
				},
			)

			// 注入环境变量
			SendMessage(versionInfo, fmt.Sprintf("注入变量: %s=%s", env.Key, val))
			if val == "" {
				continue
			}

			ldflags += fmt.Sprintf(" -X %s=%s", env.Key, val)
		}
	}

	if ldflags != "" {
		args = append(args, fmt.Sprintf("-ldflags=%s", ldflags))
	}

	if info.Tags != "" {
		args = append(args, fmt.Sprintf("-tags=%s", info.Tags))
	}

	if len(info.Flags) > 0 {
		args = append(args, info.Flags...)
	}

	args = append(args, "main.go")
	cmd := exec.Command("go", args...)
	SendMessage(versionInfo, fmt.Sprintf("编译命令: go %v", args))

	// 设置工作目录到代码目录
	cmd.Dir = workDir

	// 设置环境变量
	env := []string{
		"GOOS=" + info.Goos.String(),
		"GOARCH=" + info.Goarch.String(),
	}

	// 环境变量输出
	SendMessage(versionInfo, fmt.Sprintf("环境变量: %v", env))
	env = append(env, os.Environ()...)

	cmd.Env = env

	writer := NewWriter(versionInfo)
	cmd.Stdout = writer
	cmd.Stderr = writer

	err := cmd.Run()
	if err != nil {
		SendMessage(versionInfo, fmt.Sprintf("运行失败: %s", writer.String()))
		return err
	}

	SendMessage(versionInfo, fmt.Sprintf("运行成功: %s", writer.String()))
	return nil
}

func SendMessage(versionInfo model.VersionInfo, msg string) {
	connID := fmt.Sprintf("%d_%d", versionInfo.ID, versionInfo.ProjectId)
	websocket.GetWebSocketManager().SendMessage(connID, []byte(msg), 1)
}
