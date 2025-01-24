package service

import (
	"lime/internal/app/project/model"
	"lime/internal/common/service"
)

type Compile struct {
	service.BaseService[model.CompileInfo]
}

func NewCompile(args ...any) *Compile {
	srv := &Compile{}
	srv.BaseService = service.NewSrv(srv.BaseService, args...)
	return srv
}
