package global

import (
	"io/fs"
	"lime/pkg/logger"
	"path/filepath"

	"gorm.io/gorm"
)

var (
	MsgChannel = make(chan string, 10)
	Logger     *logger.LoggerWrapper
	DB         *gorm.DB
	S_WEB      fs.FS
)

const (
	TOPIC_CONSOLE_LOG = "TOPIC_CONSOLE_LOG"
	ROOT_PATH         = "root"
	DEFAULT_PWD       = "123456"
	DB_Key            = "data_base:info"
)

var (
	STATIC_PATH = filepath.Join(ROOT_PATH, "static")
	TMP_PATH    = filepath.Join(ROOT_PATH, "tmp")
	PKG_PATH    = filepath.Join(ROOT_PATH, "packages")
	DATA_PATH   = filepath.Join(ROOT_PATH, "data")
)

func WriteLog(msg string) {
	MsgChannel <- msg
}

var (
	APP_NAME = "青柠版本管理系统"
	VERSION  = "v1.0.0.1"
)
