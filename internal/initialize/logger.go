package initialize

import (
	"lime/internal/common/config"
	"lime/internal/global"
	"lime/pkg/logger"
	"path/filepath"
)

func InitLogger() {
	if global.Logger != nil {
		return
	}

	cfg := new(logger.Config)
	cfg.Level = config.GetParam(config.LOG, "level", 0).Int()
	path := config.GetParam(config.LOG, "path", "logs").String()
	cfg.Path = filepath.Join(global.ROOT_PATH, path)
	cfg.Name = "host_log"
	cfg.MaxSize = config.GetParam(config.LOG, "max-size", 100).Int() // MB
	cfg.MaxBackups = config.GetParam(config.LOG, "max-backups", 100).Int()
	cfg.MaxAge = config.GetParam(config.LOG, "max-age", 100).Int() // days
	cfg.Compress = config.GetParam(config.LOG, "compress", true).Bool()

	var err error
	global.Logger, err = logger.NewLoggerWrapper(*cfg)
	if err != nil {
		panic(err)
	}
}
