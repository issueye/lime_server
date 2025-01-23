package initialize

import (
	"lime/internal/common/config"
)

func InitConfig() {
	config.InitConfig()

	config.SetParamExist(config.SERVER, "port", "端口号", config.DEF_PORT)
	config.SetParamExist(config.SERVER, "mode", `服务运行模式， debug \ release`, "debug")

	config.SetParamExist(config.LOG, "path", "日志存放路径", "logs")
	config.SetParamExist(config.LOG, "max-size", "日志大小", 100)
	config.SetParamExist(config.LOG, "max-backups", "最大备份数", 10)
	config.SetParamExist(config.LOG, "max-age", "保存天数", 10)
	config.SetParamExist(config.LOG, "compress", "是否压缩", true)
	config.SetParamExist(config.LOG, "level", "日志输出等级", -1)

	// 初始化 jwt key 随机生成的码
	config.SetParamExist(config.JWT, "jwt-secret-key", "jwt 密钥", "pkkwmjjum5hvfqybnbxo97ol2spriy49")
}
