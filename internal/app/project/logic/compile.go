package logic

import (
	"fmt"
	"lime/internal/app/project/model"
	"lime/internal/app/project/requests"
	"lime/internal/app/project/service"
	"os"
	"os/exec"

	"github.com/dop251/goja"
)

// CompileProject 编译项目
func CompileProject(req *requests.CompileRequest) error {
	srv := service.NewCompile()
	err := srv.Create(&req.CompileInfo)
	if err != nil {
		return err
	}

	// 执行编译
	return CompileProgram(req.CompileInfo)
}

func CompileProgram(info model.CompileInfo) error {
	// 执行编译前的JavaScript脚本
	if err := executeScripts(info.Scripts, info); err != nil {
		return err
	}

	// 准备编译命令
	args := []string{"build", "-o", info.Output}

	// 设置编译标志
	ldflags := info.Ldflags
	if len(info.EnvVars) > 0 {
		for _, env := range info.EnvVars {
			ldflags += " -X " + fmt.Sprintf(" -X %s=%s", env.Key, env.Value)
		}
	}

	if info.Ldflags != "" {
		args = append(args, fmt.Sprintf("-ldflags=%s", info.Ldflags))
	}

	if info.Tags != "" {
		args = append(args, fmt.Sprintf("-tags=%s", info.Tags))
	}
	if len(info.Flags) > 0 {
		args = append(args, info.Flags...)
	}

	cmd := exec.Command("go", args...)

	// 设置环境变量
	env := append(os.Environ(),
		"GOOS="+info.Goos,
		"GOARCH="+info.Goarch,
	)
	cmd.Env = env

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

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
	}

	return nil
}
