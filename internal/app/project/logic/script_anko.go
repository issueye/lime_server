package logic

import (
	"context"
	"fmt"
	"lime/internal/app/project/model"
	"os"
	"os/exec"
	"strings"

	"github.com/mattn/anko/env"
	_ "github.com/mattn/anko/packages"
	"github.com/mattn/anko/vm"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Anko struct {
	ctx    context.Context
	env    *env.Env
	params map[string]any
}

func NewAnko() *Anko {
	e := env.NewEnv()
	ak := &Anko{
		ctx:    context.Background(),
		env:    e,
		params: make(map[string]any),
	}

	ak.InjectEnv()
	ak.InjectFunc()

	return ak
}

// GetEnv
func (a *Anko) GetEnv() *env.Env {
	return a.env
}

// SetParams
func (a *Anko) SetParams(name string, elem any) {
	a.params[name] = elem
	a.env.Define(name, elem)
}

// GetParams
func (a *Anko) GetParams(name string) any {
	return a.params[name]
}

func (a *Anko) InjectFunc() {
	// 注入方法
	a.env.Define("println", a.Println)       // 输出到控制台
	a.env.Define("runCommand", a.RunCommand) // 运行命令行命令
	a.env.Define("env", os.Getenv)           // 获取环境变量
	// a.env.Define("Json", NewJson)                 // 新建Json对象
	// a.env.Define("JsonFromFile", NewJsonFromFile) // 新建Json对象

	// 注入 Json 相关函数和方法
	a.env.Define("Json", NewJson)                 // 新建Json对象
	a.env.Define("JsonFromFile", NewJsonFromFile) // 从文件新建Json对象

	// 注册 Json 类型和方法
	err := a.env.DefineType("Json", &Json{})
	if err != nil {
		panic(err)
	}

	// 注入 Json 结构体的所有方法
	a.env.Define("GetString", (*Json).GetString)
	a.env.Define("GetInt", (*Json).GetInt)
	a.env.Define("GetBool", (*Json).GetBool)
	a.env.Define("GetFloat", (*Json).GetFloat)
	a.env.Define("GetStrings", (*Json).GetStrings)
	a.env.Define("GetInts", (*Json).GetInts)
	a.env.Define("GetBools", (*Json).GetBools)
	a.env.Define("GetFloats", (*Json).GetFloats)
	a.env.Define("Set", (*Json).Set)
	a.env.Define("Delete", (*Json).Delete)
	a.env.Define("ToString", (*Json).String)
	a.env.Define("Save", (*Json).Save)
}

// 注入变量
func (a *Anko) InjectEnv() {
	a.env.Define("GOOS", os.Getenv("GOOS"))     // 操作系统
	a.env.Define("GOARCH", os.Getenv("GOARCH")) // 架构
}

func (a *Anko) Execute(code string) (any, error) {
	data, err := vm.Execute(a.GetEnv(), nil, code)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (a *Anko) Println(args ...any) {
	version, ok := a.params["version"]
	if ok {
		msgs := make([]string, 0, len(args))
		for _, arg := range args {
			msgs = append(msgs, fmt.Sprint(arg))
		}

		SendMessage(version.(model.VersionInfo), strings.Join(msgs, " "))
	}

	fmt.Println(args...)
}

// 运行命令行命令
func (a *Anko) RunCommand(target string, args ...string) (string, error) {
	// 创建命令行命令
	cmd := exec.Command(target, args...)
	// 执行命令行命令

	version, ok := a.params["version"]
	if ok {
		writer := NewWriter(version.(model.VersionInfo))
		cmd.Stdout = writer
		cmd.Stderr = writer

		// 执行命令行命令
		err := cmd.Run()
		return writer.String(), err
	}

	return "", nil
}

type Json struct {
	data   string
	Getter gjson.Result
}

func NewJson(data string) *Json {
	return &Json{
		data:   data,
		Getter: gjson.Parse(data),
	}
}

func NewJsonFromFile(file string) (*Json, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return NewJson(string(data)), nil
}

func (j *Json) GetString(path string) string {
	return j.Getter.Get(path).String()
}

func (j *Json) GetInt(path string) int64 {
	return j.Getter.Get(path).Int()
}

func (j *Json) GetBool(path string) bool {
	return j.Getter.Get(path).Bool()
}

func (j *Json) GetFloat(path string) float64 {
	return j.Getter.Get(path).Float()
}

func (j *Json) GetStrings(path string) []string {
	arr := j.Getter.Get(path).Array()
	rtn := make([]string, 0, len(arr))
	for _, v := range arr {
		rtn = append(rtn, v.String())
	}

	return rtn
}

func (j *Json) GetInts(path string) []int64 {
	arr := j.Getter.Get(path).Array()
	rtn := make([]int64, 0, len(arr))
	for _, v := range arr {
		rtn = append(rtn, v.Int())
	}

	return rtn
}

func (j *Json) GetBools(path string) []bool {
	arr := j.Getter.Get(path).Array()
	rtn := make([]bool, 0, len(arr))
	for _, v := range arr {
		rtn = append(rtn, v.Bool())
	}

	return rtn
}

func (j *Json) GetFloats(path string) []float64 {
	arr := j.Getter.Get(path).Array()
	rtn := make([]float64, 0, len(arr))
	for _, v := range arr {
		rtn = append(rtn, v.Float())
	}

	return rtn
}

func (j *Json) String() string {
	return j.data
}

func (j *Json) Set(path string, value any) {
	resData, err := sjson.Set(j.data, path, value)
	if err != nil {
		return
	}

	j.data = resData
}

func (j *Json) Delete(path string) {
	sjson.Delete(j.data, path)
}

func (j *Json) Save(file string) error {
	// 覆盖写入
	return os.WriteFile(file, []byte(j.data), 0644)
}
