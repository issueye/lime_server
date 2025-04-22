package http

import (
	adminRouter "lime/internal/app/admin/router"
	"lime/internal/common/config"
	"lime/internal/common/controller"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	adminRouter adminRouter.Router
}

func MakeRouter() Router {
	return Router{
		adminRouter: adminRouter.MakeRouter(),
	}
}

func (router *Router) RegisterRouter(engine *gin.Engine) {
	v1 := engine.Group("/api/v1")
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
		router.adminRouter.Register(v1)
	}

	engine.NoRoute(func(ctx *gin.Context) {
		slog.Error("404", slog.String("path", ctx.Request.URL.Path), slog.String("method", ctx.Request.Method))
		ctl := controller.New(ctx)
		ctl.FailWithCode(http.StatusNotFound, "not found")
	})
}
