package logic

import (
	"lime/internal/app/admin/requests"
	"lime/internal/app/admin/response"
	"lime/internal/common/config"
)

func GetSystemSettings() []*response.Settings {
	list := make([]*response.Settings, 0)

	r := config.GetParam(config.SERVER, "port", "端口号")
	list = append(list, &response.Settings{Group: config.SERVER, Key: "port", Value: r.String()})

	r = config.GetParam(config.SERVER, "mode", `服务运行模式， debug \ release`)
	list = append(list, &response.Settings{Group: config.SERVER, Key: "mode", Value: r.String()})

	r = config.GetParam(config.SERVER, "client-path", "客户端路径")
	list = append(list, &response.Settings{Group: config.SERVER, Key: "client-path", Value: r.String()})

	return list
}

func SetSystemSetting(systems []*requests.Settings) {
	for _, data := range systems {
		config.SetParam(config.SERVER, data.Key, "", data.Value)
	}
}

func GetLoggerSettings() []*response.Settings {
	list := make([]*response.Settings, 0)

	r := config.GetParam(config.LOG, "path", "日志存放路径")
	list = append(list, &response.Settings{Group: config.LOG, Key: "path", Value: r.String()})

	r = config.GetParam(config.LOG, "max-size", `日志大小`)
	list = append(list, &response.Settings{Group: config.LOG, Key: "max-size", Value: r.String()})

	r = config.GetParam(config.LOG, "max-backups", "最大备份数")
	list = append(list, &response.Settings{Group: config.LOG, Key: "max-backups", Value: r.String()})

	r = config.GetParam(config.LOG, "max-age", "保存天数")
	list = append(list, &response.Settings{Group: config.LOG, Key: "max-age", Value: r.String()})

	r = config.GetParam(config.LOG, "compress", "是否压缩")
	list = append(list, &response.Settings{Group: config.LOG, Key: "compress", Value: r.String()})

	r = config.GetParam(config.LOG, "level", "日志输出等级")
	list = append(list, &response.Settings{Group: config.LOG, Key: "level", Value: r.String()})

	return list
}

func SetLoggerSetting(systems []*requests.Settings) {
	for _, data := range systems {
		config.SetParam(config.LOG, data.Key, "", data.Value)
	}
}
