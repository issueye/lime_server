package middleware

import (
	"lime/internal/common"
	"lime/internal/common/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Store struct {
	Token    string `json:"token"`
	UserID   uint   `json:"user_id"`
	RoleCode string `json:"role_code"`
}

// 使用内存存储 token，实际应用中应使用数据库或缓存
var tokenStore = make(map[string]Store)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctl := controller.New(c)
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			ctl.FailWithCode(http.StatusUnauthorized, "Authorization 不能为空")
			c.Abort()
			return
		}

		info, err := common.ParseToken(tokenString)
		if err != nil {
			ctl.FailWithCode(http.StatusUnauthorized, "无效的 token")
			c.Abort()
			return
		}

		if info.ID > 0 {
			// 如果 tokenStore 中不存在该 token，则添加到 tokenStore 中
			if _, ok := tokenStore[tokenString]; !ok {
				tokenStore[tokenString] = Store{
					Token:  tokenString,
					UserID: info.ID,
				}
			}

			c.Set("user_id", info.ID)
			c.Next()
		} else {
			ctl.FailWithCode(http.StatusUnauthorized, "无效的 token")
			c.Abort()
			return
		}
	}
}
