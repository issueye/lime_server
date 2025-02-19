package logic

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"lime/internal/app/project/model"
	"lime/internal/app/project/repo"
	"lime/internal/app/project/requests"
	"lime/internal/app/project/service"
	"lime/internal/app/websocket"
	"lime/internal/global"
	"os"
	"os/exec"
	"path/filepath"
	"time"
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
// Compiler 编译器结构体
type Compiler struct {
	projectInfo model.ProjectInfo
	versionInfo model.VersionInfo
	compileInfo model.CompileInfo
	workDir     string
	output      string
}

// sendMessage 发送消息
func (c *Compiler) sendMessage(msg string) {
	connID := fmt.Sprintf("%d_%d", c.versionInfo.ID, c.versionInfo.ProjectId)

	response := websocket.Message{
		Type:    websocket.JsonMessage,
		Content: msg,
		Time:    time.Now(),
	}

	responseMsg, err := json.Marshal(response)
	if err != nil {
		return
	}

	websocket.GetWebSocketManager().SendMessage(connID, responseMsg, 1)
}

// sendMessagef 格式化发送消息
func (c *Compiler) sendMessagef(format string, args ...any) {
	c.sendMessage(fmt.Sprintf(format, args...))
}

// prepareOutput 准备输出文件名
func (c *Compiler) prepareOutput() error {
	output, err := GetOutfileName(InvokeParams{
		Code:    c.compileInfo.Output,
		Project: c.projectInfo,
		Version: c.versionInfo,
		Compile: c.compileInfo,
	})
	if err != nil {
		return fmt.Errorf("获取输出文件名称失败: %v", err)
	}
	c.output = output
	c.sendMessagef("输出文件名称: %s", output)
	return nil
}

// NewCompiler 创建编译器实例
func NewCompiler(project model.ProjectInfo, version model.VersionInfo, compile model.CompileInfo) *Compiler {
	return &Compiler{
		projectInfo: project,
		versionInfo: version,
		compileInfo: compile,
	}
}

// CompileProject 编译项目入口函数
func CompileProject(projectId uint, versionId uint) error {
	req, err := GetConfigByProjectId(projectId)
	if err != nil {
		return err
	}

	if req == nil {
		return fmt.Errorf("编译配置不存在")
	}

	projectSrv := service.NewProject()
	projectInfo, err := projectSrv.GetByMap(map[string]any{"id": projectId})
	if err != nil {
		return err
	}

	versionSrv := service.NewVersion()
	versionInfo, err := versionSrv.GetByMap(map[string]any{"id": versionId, "project_id": projectId})
	if err != nil {
		return err
	}

	compiler := NewCompiler(*projectInfo, *versionInfo, *req)
	return compiler.Compile()
}

// Compile 执行编译流程
func (c *Compiler) Compile() error {
	c.sendMessage("获取输出文件名称")
	if err := c.prepareOutput(); err != nil {
		c.sendMessagef("准备输出文件名称失败: %v", err)
		return err
	}

	c.sendMessage("拉取代码到临时目录")
	if err := c.prepareWorkspace(); err != nil {
		c.sendMessagef("准备工作空间失败: %v", err)
		return err
	}

	c.sendMessage("执行编译前的脚本")
	if err := c.runBeforeScripts(); err != nil {
		c.sendMessagef("执行编译前脚本失败: %v", err)
		return err
	}

	c.sendMessage("执行构建命令")
	if err := c.buildProject(); err != nil {
		c.sendMessagef("执行构建命令失败: %v", err)
		return err
	}

	c.sendMessage("执行编译后的脚本")
	if err := c.runAfterScripts(); err != nil {
		c.sendMessagef("执行编译后的脚本失败: %v", err)
		return err
	}

	c.sendMessage("保存打包文件记录")
	if err := c.savePackage(); err != nil {
		c.sendMessagef("保存打包文件失败: %v", err)
		return err
	}

	c.sendMessage("编译完成")
	return nil
}

// prepareWorkspace 准备工作空间
func (c *Compiler) prepareWorkspace() error {
	// 拉取代码
	gitRepo, err := repo.PullCode(&c.projectInfo, c.versionInfo)
	if err != nil {
		return fmt.Errorf("拉取代码失败: %v", err)
	}

	c.sendMessagef("切换代码版本: %s  -hash:%s", c.versionInfo.Version, c.versionInfo.Hash)
	if err = repo.Checkout(c.versionInfo, gitRepo); err != nil {
		return fmt.Errorf("切换代码版本失败: %v", err)
	}

	// 获取工作目录
	workDir, err := gitRepo.Worktree()
	if err != nil {
		return fmt.Errorf("获取工作目录失败: %v", err)
	}
	c.workDir = workDir.Filesystem.Root()
	return nil
}

// runBeforeScripts 执行编译前脚本
func (c *Compiler) runBeforeScripts() error {
	return c.executeScripts(c.compileInfo.BeforeScripts)
}

// buildProject 执行构建
func (c *Compiler) buildProject() error {
	return c.runCommand()
}

// runAfterScripts 执行编译后脚本
func (c *Compiler) runAfterScripts() error {
	return c.executeScripts(c.compileInfo.AfterScripts)
}

// savePackage 保存打包文件
func (c *Compiler) savePackage() error {
	// 获取文件信息
	outputPath := filepath.Join(c.workDir, c.output)
	fileInfo, err := os.Stat(outputPath)
	if err != nil {
		return fmt.Errorf("获取文件信息失败: %v", err)
	}

	hash, err := calculateFileHash(outputPath)
	if err != nil {
		return fmt.Errorf("计算文件hash失败: %v", err)
	}

	// 只获取文件名称
	fileBase := filepath.Base(c.output)
	savePath := filepath.Join(global.PKG_PATH, c.versionInfo.Version)
	c.sendMessagef("保存路径: %s", savePath)
	if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
		return fmt.Errorf("创建打包目录失败: %v", err)
	}

	newPath := filepath.Join(savePath, fileBase)
	c.sendMessagef("移动文件: %s -> %s", outputPath, newPath)
	if err := os.Rename(outputPath, newPath); err != nil {
		return fmt.Errorf("移动文件失败: %v", err)
	}

	// 通过项目编码、版本号编码查询是否存在，如果存在则删除
	// 存在则删除
	c.sendMessage("删除旧的打包记录")
	if err := service.NewPackage().DeleteByFields(map[string]any{"project_id": c.projectInfo.ID, "version_id": c.versionInfo.ID}); err != nil {
		return fmt.Errorf("删除旧的打包记录失败: %v", err)
	}

	pkg := &model.PackageInfo{
		PackageBase: model.PackageBase{
			ProjectName: c.projectInfo.Name,
			ProjectId:   c.projectInfo.ID,
			Version:     c.versionInfo.Version,
			VersionId:   c.versionInfo.ID,
			Name:        fileBase,
			Path:        newPath,
			Size:        fileInfo.Size(),
			Hash:        hash,
		},
	}

	if err := service.NewPackage().Create(pkg); err != nil {
		return fmt.Errorf("保存打包记录失败: %v", err)
	}

	c.sendMessage("打包文件记录保存完成")
	return nil
}

// executeScripts 执行编译前的JavaScript脚本
func (c *Compiler) executeScripts(scripts model.Scripts) error {
	for _, script := range scripts {
		c.sendMessage(fmt.Sprintf("执行脚本: %s", script.Name))
		err := BeforeScript(
			InvokeParams{
				WorkDir: c.workDir,
				Code:    script.Content,
				Project: c.projectInfo,
				Version: c.versionInfo,
				Compile: c.compileInfo,
			},
		)

		if err != nil {
			return err
		}
	}

	return nil
}

// runCommand 执行构建命令
func (c *Compiler) runCommand() error {
	// 准备编译命令
	args := []string{"build", "-o", c.output}

	// 设置编译标志
	ldflags := c.compileInfo.Ldflags
	if len(c.compileInfo.EnvVars) > 0 {
		for _, env := range c.compileInfo.EnvVars {
			if env.Key == "" || env.Value == "" {
				continue
			}

			// 获取环境变量的值
			val, _ := InjectEnv(
				InvokeParams{
					Code:    env.Value,
					Project: c.projectInfo,
					Version: c.versionInfo,
					Compile: c.compileInfo,
				},
			)

			// 注入环境变量
			c.sendMessagef("注入变量: %s=%s", env.Key, val)
			if val == "" {
				continue
			}

			ldflags += fmt.Sprintf(" -X %s=%s", env.Key, val)
		}
	}

	if ldflags != "" {
		args = append(args, fmt.Sprintf("-ldflags=%s", ldflags))
	}

	if c.compileInfo.Tags != "" {
		args = append(args, fmt.Sprintf("-tags=%s", c.compileInfo.Tags))
	}

	if len(c.compileInfo.Flags) > 0 {
		args = append(args, c.compileInfo.Flags...)
	}

	args = append(args, c.compileInfo.MainPath)
	cmd := exec.Command("go", args...)
	c.sendMessagef("编译命令: go %v", args)

	// 设置工作目录到代码目录
	cmd.Dir = c.workDir

	// 设置环境变量
	env := []string{
		"GOOS=" + c.compileInfo.Goos.String(),
		"GOARCH=" + c.compileInfo.Goarch.String(),
	}

	// 注入运行期环境变量
	if len(c.compileInfo.OsEnvVars) > 0 {
		for _, data := range c.compileInfo.OsEnvVars {
			if data.Key == "" || data.Value == "" {
				continue
			}

			env = append(env, fmt.Sprintf("%s=%s", data.Key, data.Value))
		}
	}

	// 环境变量输出
	c.sendMessagef("环境变量: %v", env)
	env = append(env, os.Environ()...)

	cmd.Env = env

	writer := NewWriter(c.versionInfo)
	cmd.Stdout = writer
	cmd.Stderr = writer

	err := cmd.Run()
	if err != nil {
		c.sendMessagef("运行失败: %s", writer.String())
		return err
	}

	c.sendMessagef("运行成功: %s", writer.String())
	return nil
}

// 删除原有的全局函数
func calculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
