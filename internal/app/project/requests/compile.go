package requests

import "lime/internal/app/project/model"

// CompileRequest 编译请求参数
type CompileRequest struct {
	model.CompileInfo
}

// NewCompileRequest 创建编译请求
func NewCompileRequest() *CompileRequest {
	return &CompileRequest{
		CompileInfo: model.CompileInfo{},
	}
}

// Validate 验证参数
func (r *CompileRequest) Validate() error {
	return nil
}

type SaveCompileConfigRequest struct {
	model.CompileInfo
}

func NewSaveCompileConfigRequest() *SaveCompileConfigRequest {
	return &SaveCompileConfigRequest{
		CompileInfo: model.CompileInfo{},
	}
}
