package middleware

import (
	"fmt"
	"lime/internal/app/admin/service"
	"lime/internal/common"
	"lime/internal/common/controller"
	"lime/internal/global"
	"strings"

	"github.com/gin-gonic/gin"
)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("Authorization")
		info, err := common.ParseToken(authToken)
		if err != nil {
			controller.New(c).Unauthorized("无效的 token")
			c.Abort()
			return
		}

		//获取请求的PATH
		path := c.FullPath()
		fmt.Println("path:", path)
		obj := strings.TrimPrefix(path, "/api/v1")
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := info.RoleCode
		e := service.NewCasbin(global.DB).Casbin() // 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if !success {
			msg := fmt.Sprintf("%s 无权访问 %s", sub, obj)
			controller.New(c).Forbidden(msg)
			c.Abort()
			return
		}
		c.Next()
	}
}
