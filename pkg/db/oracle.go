package db

import (
	"fmt"
	"net/url"
	"time"

	oracle "github.com/godoes/gorm-oracle"
	"go.uber.org/zap"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

func InitOracle(cfg *Config, log *zap.SugaredLogger) *gorm.DB {
	// 进行url 编码
	pwd := url.QueryEscape(cfg.Password)

	dsn := fmt.Sprintf(
		"oracle://%s:%s@%s:%d/%s",
		cfg.Username,
		pwd,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	// 隐藏密码
	showDsn := fmt.Sprintf(
		"oracle://%s:***********@%s:%d/%s",
		cfg.Username,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	newLogger := glogger.New(
		Writer{
			log: log,
		},
		glogger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  glogger.Info,           // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)

	var l glogger.Interface

	// log = logger.NewWithZap(Log, logConf)
	if cfg.LogMode {
		l = newLogger.LogMode(glogger.Info)
	} else {
		l = newLogger.LogMode(glogger.Silent)
	}

	db, err := gorm.Open(oracle.Open(dsn), &gorm.Config{
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   l,
	})

	if err != nil {
		log.Panicf("初始化oracle数据库异常: %v", err)
		panic(fmt.Errorf("初始化oracle数据库异常: %v", err))
	}

	db.Debug()
	log.Infof("初始化oracle数据库完成! dsn: %s", showDsn)

	return db
}
