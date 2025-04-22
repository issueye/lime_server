package initialize

import (
	"context"
	"lime/internal/common/config"
	"lime/internal/initialize/http"
)

type Server struct {
	Ctx        context.Context
	HttpServer *http.HttpServer
}

func NewServer(ctx context.Context) *Server {
	return &Server{
		Ctx: ctx,
	}
}

func (server *Server) Run() {
	InitRuntime()
	InitConfig()
	InitLogger()
	InitDB()

	port := config.GetParam(config.SERVER, "http-port", config.DEF_PORT).Int()
	mode := config.GetParam(config.SERVER, "mode", "debug").String()
	server.HttpServer = http.NewHttpServer(server.Ctx, port, mode)
}

func (server *Server) Stop() {
	// http 服务
	server.HttpServer.Stop()
}
