package middleware

import (
	"errors"
	"lime/internal/common/config"
	"lime/internal/common/controller"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctl := controller.New(c)
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			ctl.FailWithCode(http.StatusUnauthorized, "Authorization 不能为空")
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("错误的签名方法")
			}
			key := config.GetParam(config.JWT, "jwt-secret-key", "pkkwmjjum5hvfqybnbxo97ol2spriy49").String()
			return []byte(key), nil
		})
		if err != nil {
			ctl.FailWithCode(http.StatusUnauthorized, "无效的 token")
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 可以从 claims 中获取用户相关信息，例如用户 ID，用于后续的权限验证等操作
			userID := uint(claims["id"].(float64))
			c.Set("user_id", userID)
			c.Next()
		} else {
			ctl.FailWithCode(http.StatusUnauthorized, "无效的 token")
			c.Abort()
			return
		}
	}
}
