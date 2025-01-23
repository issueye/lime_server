package initialize

import (
	"context"
	"fmt"
	"lime/docs"
	"lime/internal/common/config"
	"lime/internal/common/middleware"
	route "lime/internal/common/router"
	"lime/internal/global"
	"net/http"
	"time"

	"github.com/TelenLiu/knife4j_go"
	"github.com/gin-gonic/gin"
)

func InitHttpServer(ctx context.Context) {
	port := config.GetParam(config.SERVER, "http-port", config.DEF_PORT).Int()
	mode := config.GetParam(config.SERVER, "mode", "debug").String()
	gin.SetMode(mode)
	// gin引擎对象
	engine := gin.New()

	// 中间件
	engine.Use(middleware.Cors())
	engine.Use(middleware.Logger())
	engine.Use(middleware.Recovery())
	route.InitRouter(engine)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: engine,
	}

	knife4j_go.SetDiyContent("doc.json", []byte(docs.SwaggerInfo.ReadDoc()))
	engine.StaticFS("/doc", http.FS(knife4j_go.GetKnife4jVueDistRoot()))
	engine.GET("/services.json", func(c *gin.Context) {
		c.String(200, `[
		    {
				"name": "定时任务调度服务系统v1.0",
				"url": "/doc.json",
				"swaggerVersion": "2.0",
				"location": "/doc.json",
			}
		]`)
	})

	engine.StaticFS("/web", http.FS(global.S_WEB))

	global.HttpServer = srv
	global.HttpEngine = engine

	go func(_ context.Context) {
		if err := global.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err.Error())
		}

		global.WriteLog("HTTP服务关闭 --->")
	}(ctx)
}

func FreeHttpServer() {
	if global.HttpServer == nil {
		return
	}

	global.WriteLog("HTTP服务关闭")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := global.HttpServer.Shutdown(ctx)
	if err != nil {
		global.WriteLog(fmt.Sprintf("HTTP服务关闭失败 %s", err.Error()))
	} else {
		global.WriteLog("HTTP服务关闭成功")
	}
}
