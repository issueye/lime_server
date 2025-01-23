package route

import (
	adminRouter "lime/internal/app/admin/router"
	projectRouter "lime/internal/app/project/router"
	"lime/internal/common/config"
	"lime/internal/common/controller"
	"lime/internal/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctl := controller.New(ctx)
			ctl.SuccessData(map[string]any{"msg": "pong"})
		})

		v1.GET("/version", func(ctx *gin.Context) {
			ctl := controller.New(ctx)
			versionInfo := config.GetVersionInfo()
			ctl.SuccessData(map[string]any{
				"version":     versionInfo.Version,
				"build_time":  versionInfo.BuildTime,
				"git_commit":  versionInfo.GitCommit,
				"environment": versionInfo.Environment,
			})
		})

		// 注册管理路由
		adminRouter.Register(v1)
		// 注册项目管理路由
		projectRouter.Register(v1)
	}

	r.NoRoute(func(ctx *gin.Context) {
		global.Logger.Logger.Error("404", zap.String("path", ctx.Request.URL.Path), zap.String("method", ctx.Request.Method))
		ctl := controller.New(ctx)
		ctl.FailWithCode(http.StatusNotFound, "not found")
	})
}
