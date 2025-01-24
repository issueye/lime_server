package service

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dop251/goja"
)

type Script struct {
	Name    string `json:"name"`    // 脚本名称
	Content string `json:"content"` // 脚本内容
}

type EnvVar struct {
	Key   string `json:"key"`   // 环境变量键
	Value string `json:"value"` // 环境变量值
}

// CompileInfo holds the information for the compilation process
type CompileInfo struct {
	Output  string   `json:"output"`   // 输出文件名
	Goos    string   `json:"goos"`     // 目标操作系统
	Goarch  string   `json:"goarch"`   // 目标架构
	Flags   []string `json:"flags"`    // 编译标志
	Ldflags string   `json:"ldflags"`  // 链接标志
	Tags    string   `json:"tags"`     // 编译标签
	Scripts []Script `json:"scripts"`  // 编译前执行的JavaScript脚本
	EnvVars []EnvVar `json:"env_vars"` // 环境变量
}

// executeScripts 执行编译前的JavaScript脚本
func executeScripts(scripts []Script, info CompileInfo) error {
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

// SaveCompileInfo saves the compile information to a file
func SaveCompileInfo(info CompileInfo) error {
	file, err := os.Create("compile_info.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf(
		"Output: %s\nGoos: %s\nGoarch: %s\nFlags: %v\nLdflags: %s\nTags: %s\nScripts: %v\nEnvVars: %v\n",
		info.Output, info.Goos, info.Goarch, info.Flags, info.Ldflags, info.Tags, info.Scripts, info.EnvVars,
	))
	return err
}

// CompileProgram compiles the Go program based on the provided compile information
func CompileProgram(info CompileInfo) error {
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
