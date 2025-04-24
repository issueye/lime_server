package initialize

import (
	"lime/internal/app/admin/logic"
	"lime/internal/app/admin/model"
	"lime/internal/common"
)

func InitDictData() {
	dicts := []model.DictsBase{
		{
			Code:        "gender",
			Name:        "性别",
			ContentType: common.CTT_INT,
			Description: "性别",
			Details: []model.DictDetail{
				{Key: "0", Value: "男", Description: "男", Extra: ""},
				{Key: "1", Value: "女", Description: "女", Extra: ""},
				{Key: "2", Value: "未知", Description: "未知", Extra: ""},
			},
		},
		{
			Code: "api_group",
			Name: "API分组",
			Details: []model.DictDetail{
				{Key: "sys", Value: "系统管理", Description: "系统管理", Extra: ""},
				{Key: "user", Value: "用户管理", Description: "用户管理", Extra: ""},
			},
		},
		{
			Code: "api_method",
			Name: "API方法",
			Details: []model.DictDetail{
				{Key: "GET", Value: "GET", Description: "GET", Extra: ""},
				{Key: "POST", Value: "POST", Description: "POST", Extra: ""},
				{Key: "PUT", Value: "PUT", Description: "PUT", Extra: ""},
				{Key: "DELETE", Value: "DELETE", Description: "DELETE", Extra: ""},
				{Key: "OPTIONS", Value: "OPTIONS", Description: "OPTIONS", Extra: ""},
				{Key: "HEAD", Value: "HEAD", Description: "HEAD", Extra: ""},
			},
		},
		{
			Code: "http_status",
			Name: "HTTP状态码",
			Details: []model.DictDetail{
				{Key: "200", Value: "OK", Description: "成功", Extra: ""},
				{Key: "400", Value: "Bad Request", Description: "错误请求", Extra: ""},
				{Key: "401", Value: "Unauthorized", Description: "未授权", Extra: ""},
				{Key: "403", Value: "Forbidden", Description: "禁止访问", Extra: ""},
				{Key: "404", Value: "Not Found", Description: "未找到", Extra: ""},
				{Key: "500", Value: "Internal Server Error", Description: "内部服务器错误", Extra: ""},
				{Key: "502", Value: "Bad Gateway", Description: "错误网关", Extra: ""},
				{Key: "503", Value: "Service Unavailable", Description: "服务不可用", Extra: ""},
				{Key: "504", Value: "Gateway Timeout", Description: "网关超时", Extra: ""},
			},
		},
	}

	lc := logic.NewDictsLogic()
	for _, dict := range dicts {
		data := model.NewDictsInfo(dict)
		lc.IsNotExistAdd(*data)
	}
}
