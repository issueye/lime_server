package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	Path       string
	Name       string     // 日志名称
	MaxSize    int        //文件大小限制,单位MB
	MaxAge     int        //日志文件保留天数
	MaxBackups int        //最大保留日志文件数量
	Compress   bool       //是否压缩处理
	Level      int        // 等级
	Mode       LogOutMode // 日志输出模式
}

type LoggerWrapper struct {
	*zap.Logger
	lumberjack *lumberjack.Logger
	path       string
}

func NewLoggerWrapper(cfg Config) (*LoggerWrapper, error) {

	path := fmt.Sprintf("%s/%s.log", cfg.Path, cfg.Name)
	fmt.Println("日志路径:", path)
	lumberjackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	writeSyncer := zapcore.AddSync(lumberjackLogger)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.Level(cfg.Level))

	var core zapcore.Core
	if cfg.Level == -1 {
		core = zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout)), atomicLevel)
	} else {
		core = zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(writeSyncer), atomicLevel)
	}

	logger := zap.New(core)

	dir := filepath.Dir(lumberjackLogger.Filename)

	return &LoggerWrapper{
		Logger:     logger,
		lumberjack: lumberjackLogger,
		path:       dir,
	}, nil
}

func (l *LoggerWrapper) Sync() {
	l.Logger.Sync()
	l.lumberjack.Rotate()
}

func (l *LoggerWrapper) Close() {
	l.Logger.Sync()

	l.NewFilename(fmt.Sprintf("%s/temp.log", l.path))

	// l.lumberjack.Rotate()
	l.lumberjack.Close()
}

func (l *LoggerWrapper) NewFilename(filename string) {
	l.lumberjack.Filename = filename
}
