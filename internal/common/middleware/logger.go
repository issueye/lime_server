package middleware

import (
	"lime/internal/global"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return ginzap.GinzapWithConfig(global.Logger, &ginzap.Config{
		TimeFormat: "2006-01-02 15:04:05",
		UTC:        true,
		SkipPaths:  []string{"/health"},
	})
}

func Recovery() gin.HandlerFunc {
	return ginzap.RecoveryWithZap(global.Logger, true)
}
