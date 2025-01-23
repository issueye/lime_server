package config

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"lime/pkg/db"
	"lime/pkg/logger"
	"lime/pkg/utils"

	"gorm.io/gorm"
)

const (
	Version     = "1.0.0" // 当前版本号
	BuildTime   = ""      // 编译时间
	GitCommit   = ""      // Git提交ID
	Environment = "dev"   // 运行环境
)

var configDB *gorm.DB

type Result struct {
	Param
}

func (r *Result) String() string {
	return r.Value.String()
}

func (r *Result) Int64() int64 {
	i, err := strconv.ParseInt(r.Value.String(), 10, 64)
	if err != nil {
		return 0
	}

	return i
}

func (r *Result) Int() int {
	i, err := strconv.Atoi(r.Value.String())
	if err != nil {
		return 0
	}

	return i
}

func (r *Result) Float64() float64 {
	i, err := strconv.ParseFloat(r.Value.String(), 64)
	if err != nil {
		return 0
	}

	return i
}

func (r *Result) Bool() bool {
	i, err := strconv.ParseBool(r.Value.String())
	if err != nil {
		return false
	}

	return i
}

func (r *Result) Datetime() *time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", r.Value.String(), time.Local)
	if err != nil {
		return nil
	}

	return &t
}

func (r *Result) ToJson() string {
	name, err := json.Marshal(r)
	if err != nil {
		return ""
	}

	return string(name)
}

func (r *Result) Description() string {
	return r.Mark
}

func SetParamExist(group, name, mark string, value any) *Result {
	r := GetParam(group, name, value)
	if r.ID == 0 {
		r = SetParam(group, name, mark, value)
	}

	return r
}

func SetParam(group, name, mark string, value any) *Result {
	data := NewData(value)

	r := GetParam(group, name, "")
	if r.ID == 0 {
		r.ID = utils.GenID()
		r.Group = strings.ToLower(group)
		r.Name = strings.ToLower(name)
		r.Value = data
		r.Mark = mark

		err := GetDB().Model(r).Create(r).Error
		if err != nil {
			return nil
		}
	} else {
		r.Group = strings.ToLower(group)
		r.Name = strings.ToLower(name)
		r.Value = data
		// r.Mark = mark
		err := GetDB().Model(r).Where("id = ?", r.ID).Updates(r).Error
		if err != nil {
			return nil
		}
	}

	return r
}

func GetParam(group, name string, DefValue any) *Result {
	r := new(Result)
	err := GetDB().Model(r).Where(`"group" = ? and name = ?`, strings.ToLower(group), strings.ToLower(name)).Find(r).Error
	if err != nil {
		r.ID = 0
		r.Name = name
		r.Value = NewData(DefValue)
		r.Mark = ""
	}

	if r.ID == 0 {
		r.Name = name
		r.Value = NewData(DefValue)
	}

	return r
}

type VersionInfo struct {
	Version     string `json:"version"`
	BuildTime   string `json:"build_time"`
	GitCommit   string `json:"git_commit"`
	Environment string `json:"environment"`
}

func GetVersionInfo() VersionInfo {
	return VersionInfo{
		Version:     Version,
		BuildTime:   BuildTime,
		GitCommit:   GitCommit,
		Environment: Environment,
	}
}

func GetDB() *gorm.DB {
	if configDB == nil {
		InitConfig()
	}

	return configDB
}

func InitConfig() {
	if configDB != nil {
		return
	}

	// 检查文件是否存在
	path := filepath.Join("root", "config", "config.db")
	logPath := filepath.Join("root", "logs", "config.log")
	log, _, err := logger.NewZap(logPath, logger.LOM_DEBUG)
	if err != nil {
		panic(fmt.Errorf("创建日志对象失败，失败原因：%s", err))
	}

	configDB = db.InitSqlite(path, log.Sugar())

	// 初始化数据库表
	// 创建表
	configDB.AutoMigrate(
		&Param{},
	)
}
