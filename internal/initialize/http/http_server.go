package http

import (
	"context"
	"fmt"
	"lime/docs"
	"lime/internal/global"
	"lime/internal/middleware"
	"net/http"
	"time"

	"github.com/TelenLiu/knife4j_go"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Router Router
	Http   *http.Server
	engine *gin.Engine
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		engine: gin.New(),
		Router: MakeRouter(),
	}
}

func (server *HttpServer) Run(ctx context.Context, port int, mode string) {
	gin.SetMode(mode)

	// 中间件
	server.engine.Use(middleware.Cors())
	server.engine.Use(middleware.Logger())
	server.engine.Use(middleware.Recovery())

	// 路由注册
	server.Router.RegisterRouter(server.engine)

	server.Http = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: server.engine,
	}

	knife4j_go.SetDiyContent("doc.json", []byte(docs.SwaggerInfo.ReadDoc()))
	server.engine.StaticFS("/doc", http.FS(knife4j_go.GetKnife4jVueDistRoot()))
	server.engine.GET("/services.json", func(c *gin.Context) {
		c.String(200, `[
		    {
				"name": "定时任务调度服务系统v1.0",
				"url": "/doc.json",
				"swaggerVersion": "2.0",
				"location": "/doc.json",
			}
		]`)
	})

	server.engine.StaticFS("/web", http.FS(global.S_WEB))

	go func(_ context.Context) {
		if err := server.Http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err.Error())
		}

		global.WriteLog("HTTP服务关闭 --->")
	}(ctx)
}

func (server *HttpServer) Stop() {
	if server.Http == nil {
		return
	}

	global.WriteLog("HTTP服务关闭")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Http.Shutdown(ctx)
	if err != nil {
		global.WriteLog(fmt.Sprintf("HTTP服务关闭失败 %s", err.Error()))
	} else {
		global.WriteLog("HTTP服务关闭成功")
	}
}
