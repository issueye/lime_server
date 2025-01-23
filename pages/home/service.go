package home

import (
	"fmt"
	"lime/internal/common/config"
	"lime/internal/global"
	"os"
	"os/exec"
	"runtime"
)

// 初始化数据
func (f *TFrmHome) InitData() {
	f.ShowRunInfo()
}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		return fmt.Errorf("unsupported operating system")
	}
	return cmd.Start()
}

func (f *TFrmHome) ShowRunInfo() {
	f.Lbl_name.SetCaption("名称：" + global.APP_NAME)
	port := config.GetParam(config.SERVER, "http-port", config.DEF_PORT).Int()
	f.Lbl_port.SetCaption(fmt.Sprintf("端口：%d", port))
	f.Lbl_version.SetCaption("版本：" + global.VERSION)

	item_pid := f.StatusBar.Panels().Items(0)
	item_pid.SetText(fmt.Sprintf("PID：%d", os.Getpid()))
}
