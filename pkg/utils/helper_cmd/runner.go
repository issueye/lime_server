package helper_cmd

import (
	"context"
	"io"
	"lime/internal/global"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

type RunResult struct {
	Msg     chan string
	Pid     int
	Process *os.Process
	Ctx     context.Context
	Cancel  context.CancelFunc
}

type RunFunc func(isSetWorkSpace bool, path string, args ...string) (*RunResult, error)

func Run(chanLen int) RunFunc {
	return func(isSetWorkSpace bool, path string, args ...string) (*RunResult, error) {

		rtn := &RunResult{}
		rtn.Ctx, rtn.Cancel = context.WithCancel(context.Background())

		cmd := exec.CommandContext(rtn.Ctx, path, args...)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

		// 获取当前程序的所在路径
		pwd, _ := os.Getwd()

		// 获取当前操作系统的环境变量
		env := os.Environ()
		// 拼接新的环境变量，将 bin 目录添加到 PATH 中
		newPath := filepath.Join(pwd, global.ROOT_PATH, "bin")
		env = append(env, "PATH="+newPath+";"+os.Getenv("PATH"))
		// fmt.Println("env", strings.Join(env, "\n"))
		// 设置环境变量
		cmd.Env = env

		// 使用传入进来的path 设置为工作区
		if isSetWorkSpace {
			workSpace := path
			// 判断是文件还是文件夹，如果是文件，则获取其父目录
			if info, err := os.Stat(workSpace); err == nil && !info.IsDir() {
				workSpace = filepath.Dir(workSpace)
			}

			cmd.Dir = workSpace
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			return nil, err
		}

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return nil, err
		}

		err = cmd.Start()
		if err != nil {
			return nil, err
		}

		if chanLen < 20 {
			chanLen = 20
		}

		rtn.Pid = cmd.Process.Pid
		rtn.Process = cmd.Process
		rtn.Msg = make(chan string, chanLen)

		go io.Copy(&appLogWriter{msgChannel: rtn.Msg}, stderr)
		go io.Copy(&appLogWriter{msgChannel: rtn.Msg}, stdout)

		return rtn, nil
	}
}
