package middleware

import (
	"lime/internal/app/admin/service"
	"lime/internal/common"
	"lime/internal/common/controller"
	"lime/internal/global"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		info, err := common.ParseToken(c.GetHeader("Authorization"))
		if err != nil {
			controller.New(c).FailWithCode(401, "无效的 token")
			c.Abort()
			return
		}

		//获取请求的PATH
		path := c.Request.URL.Path
		obj := strings.TrimPrefix(path, "/api/v1")
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(info.ID))
		e := service.NewCasbin(global.DB).Casbin() // 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if !success {
			controller.New(c).Forbidden("权限不足")
			c.Abort()
			return
		}
		c.Next()
	}
}
