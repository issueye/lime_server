package utils

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// getChildPIDs 获取指定 PID 的所有子进程 PID
func getChildPIDs(pid int) ([]int, error) {
	var childPIDs []int
	snapshot, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPALL, 0)
	if err != nil {
		return nil, err
	}
	defer windows.CloseHandle(snapshot)

	var entry windows.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))

	// 遍历进程快照
	err = windows.Process32First(snapshot, &entry)
	for err == nil {
		if entry.ParentProcessID == uint32(pid) {
			childPIDs = append(childPIDs, int(entry.ProcessID))
		}
		err = windows.Process32Next(snapshot, &entry)
	}

	return childPIDs, nil
}

type Process struct {
	Msg chan string
}

func (p *Process) sendMessage(msg string) {
	if p.Msg != nil {
		p.Msg <- msg
	}
}

// killProcessAndChildren 递归杀死指定 PID 的进程及其所有子进程
func (p *Process) KillProcessAndChildren(pid int) error {
	childPIDs, err := getChildPIDs(pid)
	if err != nil {
		return err
	}

	// 杀死子进程
	for _, childPID := range childPIDs {
		p.sendMessage(fmt.Sprintf("开始结束子进程: %d\n", childPID))

		// 打开子进程句柄
		childHandle, err := windows.OpenProcess(windows.PROCESS_TERMINATE, false, uint32(childPID))
		if err != nil {
			p.sendMessage(fmt.Sprintf("打开子进程句柄失败 %d: %v\n", childPID, err))
			continue // 继续尝试下一个子进程
		}

		// 递归杀死子进程的子进程
		if err := p.KillProcessAndChildren(childPID); err != nil {
			p.sendMessage(fmt.Sprintf("结束子进程失败 %d: %v\n", childPID, err))
		}

		// 终止子进程
		err = syscall.TerminateProcess(syscall.Handle(childHandle), 1) // 类型转换
		if err != nil {
			p.sendMessage(fmt.Sprintf("结束子进程失败 %d: %v\n", childPID, err))
		}
		windows.CloseHandle(childHandle) // 关闭句柄

		p.sendMessage(fmt.Sprintf("结束子进程成功 %d\n", childPID))
	}

	// 杀死父进程
	p.sendMessage(fmt.Sprintf("结束父进程: %d\n", pid))
	parentHandle, err := windows.OpenProcess(windows.PROCESS_TERMINATE, false, uint32(pid))
	if err != nil {
		p.sendMessage(fmt.Sprintf("打开父进程句柄失败 %d: %v\n", pid, err))
		return err
	}
	defer windows.CloseHandle(parentHandle)

	err = syscall.TerminateProcess(syscall.Handle(parentHandle), 1) // 类型转换
	if err != nil {
		p.sendMessage(fmt.Sprintf("结束父进程失败 %d: %v\n", pid, err))
		return err
	}

	p.sendMessage(fmt.Sprintf("结束父进程成功 %d\n", pid))
	return nil
}
