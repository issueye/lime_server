package logic

import (
	"context"
	"fmt"
	"lime/internal/app/project/model"

	"github.com/dop251/goja"
	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

type VM struct {
	ctx         context.Context
	VM          *goja.Runtime
	Params      map[string]any
	CodeContent string
}

type Fn func(vm *goja.Runtime, args ...any) error

func NewVM(ctx context.Context) *VM {
	vm := &VM{
		ctx: ctx,
		VM:  goja.New(),
	}
	vm.Params = make(map[string]any)
	vm.VM.Set("ctx", ctx)
	vm.injectNative()
	return vm
}

func (v *VM) ExoprtFunc(name string, fn Fn) (goja.Callable, error) {
	if fn != nil {
		err := fn(v.VM)
		if err != nil {
			return nil, err
		}
	}

	_, err := v.VM.RunString(v.CodeContent)
	if err != nil {
		return nil, err
	}

	callFn, ok := goja.AssertFunction(v.VM.Get(name))
	if !ok {
		return nil, fmt.Errorf("function %s not found", name)
	}

	return callFn, nil
}

func (v *VM) injectNative() {
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

	v.VM.Set("console", console)
}

func GetOutfileName(projectInfo model.ProjectInfo, versionInfo model.VersionInfo, info model.CompileInfo) (string, error) {
	if info.Output == "" {
		return "", fmt.Errorf("输出文件名称不能为空")
	}

	vm := NewVM(context.Background())
	vm.Params = map[string]interface{}{
		"info":    info,
		"project": projectInfo,
		"version": versionInfo,
	}

	vm.CodeContent = info.Output

	callFn, err := vm.ExoprtFunc("getOutfileName", nil)
	if err != nil {
		return "", err
	}

	res, err := callFn(goja.Undefined(),
		vm.VM.ToValue(info),
		vm.VM.ToValue(projectInfo),
		vm.VM.ToValue(versionInfo),
	)
	if err != nil {
		return "", err
	}

	return res.String(), nil
}

func NewVM2(info model.CompileInfo) (interface{}, error) {
	e := env.NewEnv()
	e.Define("println", fmt.Println)
	e.Define("data", info)

	script := `
	println("data -> ", data)
	data.Output
	`
	data, err := vm.Execute(e, nil, script)
	if err != nil {
		return nil, err
	}

	return data, nil
}
