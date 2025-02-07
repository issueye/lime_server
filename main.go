package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"lime/internal/common/config"
	"lime/internal/global"
	"lime/internal/initialize"
	"lime/pages/home"
	_ "lime/winappres"
	"os"

	"github.com/ying32/govcl/vcl"
)

//	@title			青柠版本管理系统服务v1.0
//	@version		V1.1
//	@description	青柠版本管理系统服务v1.0

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//go:embed web/*
var webDir embed.FS

var (
	VERSION = "v1.0.0"
)

func showVersion() {
	versionInfo := config.GetVersionInfo()
	fmt.Printf("Version: %s\n", VERSION)
	fmt.Printf("Build Time: %s\n", versionInfo.BuildTime)
	fmt.Printf("Git Commit: %s\n", versionInfo.GitCommit)
	fmt.Printf("Environment: %s\n", versionInfo.Environment)
	os.Exit(0)
}

func main() {
	// 处理命令行参数
	versionFlag := flag.Bool("v", false, "显示版本信息")
	flag.Parse()

	if *versionFlag {
		showVersion()
	}

	staticFp, _ := fs.Sub(webDir, "web")
	global.S_WEB = staticFp

	// 创建 context 用于管理服务生命周期
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 初始化并启动服务
	initialize.RunServer(ctx)

	// 运行 GUI 应用
	vcl.RunApp(&home.FrmHome)

	// 关闭服务
	initialize.StopServer()
}
