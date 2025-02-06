package logic

import (
	"context"
	"fmt"

	"github.com/dop251/goja"
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
