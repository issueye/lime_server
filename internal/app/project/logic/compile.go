package logic

import (
	"fmt"
	"lime/internal/app/project/model"
	"lime/internal/app/project/repo"
	"lime/internal/app/project/requests"
	"lime/internal/app/project/service"
	"os"
	"os/exec"

	"github.com/dop251/goja"
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
	output, err := GetOutfileName(projectInfo, versionInfo, info)
	if err != nil {
		return fmt.Errorf("获取输出文件名称失败: %v", err)
	}

	fmt.Println("输出路径:", output)

	// 拉取代码到临时目录
	gitRepo, err := repo.PullCode(&projectInfo, versionInfo)
	if err != nil {
		return fmt.Errorf("拉取代码失败: %v", err)
	}

	// checkout 到指定版本
	err = repo.Checkout(versionInfo, gitRepo)
	if err != nil {
		return fmt.Errorf("切换代码版本失败: %v", err)
	}

	// 获取仓库工作目录
	workDir, err := gitRepo.Worktree()
	if err != nil {
		return fmt.Errorf("获取工作目录失败: %v", err)
	}

	// 执行编译前的JavaScript脚本
	if err := executeScripts(info.Scripts, info); err != nil {
		return fmt.Errorf("执行脚本失败: %v", err)
	}

	// 准备编译命令
	args := []string{"build", "-o", output}

	// 设置编译标志
	ldflags := info.Ldflags
	if len(info.EnvVars) > 0 {
		for _, env := range info.EnvVars {
			ldflags += fmt.Sprintf(" -X %s=%s", env.Key, env.Value)
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
	fmt.Println("编译命令: go", args)

	// 设置工作目录到代码目录
	cmd.Dir = workDir.Filesystem.Root()

	// 设置环境变量
	env := append(os.Environ(),
		"GOOS="+info.Goos.String(),
		"GOARCH="+info.Goarch.String(),
	)

	cmd.Env = env
	// 环境变量输出
	fmt.Println("环境变量:", env)

	writer := NewWriter()
	cmd.Stdout = writer
	cmd.Stderr = writer

	return cmd.Run()
}

// executeScripts 执行编译前的JavaScript脚本
func executeScripts(scripts []model.Script, info model.CompileInfo) error {
	vm := goja.New()

	// 注入编译信息到JS环境
	compileInfo := map[string]interface{}{
		"output":  info.Output,
		"goos":    info.Goos,
		"goarch":  info.Goarch,
		"flags":   info.Flags,
		"ldflags": info.Ldflags,
		"tags":    info.Tags,
		"envVars": info.EnvVars,
	}
	err := vm.Set("compileInfo", compileInfo)
	if err != nil {
		return fmt.Errorf("设置编译信息到JS环境失败: %v", err)
	}

	// 注入console.log
	console := map[string]interface{}{
		"log": func(call goja.FunctionCall) goja.Value {
			args := make([]interface{}, len(call.Arguments))
			for i, arg := range call.Arguments {
				args[i] = arg.Export()
			}
			fmt.Println(args...)
			return goja.Undefined()
		},
	}

	err = vm.Set("console", console)
	if err != nil {
		return fmt.Errorf("设置console到JS环境失败: %v", err)
	}

	// 执行每个脚本
	for i, script := range scripts {
		fmt.Printf("执行脚本 #%d: %s\n", i+1, script.Name)
		_, err := vm.RunString(script.Content)
		if err != nil {
			return fmt.Errorf("执行脚本 #%d 失败: %v", i+1, err)
		}

		fmt.Printf("脚本 #%d 执行完毕\n", i+1)
	}

	return nil
}
