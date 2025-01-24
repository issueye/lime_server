package requests

// CompileRequest 编译请求参数
type CompileRequest struct {
	Output  string   `json:"output" binding:"required"` // 输出文件名
	Goos    string   `json:"goos" binding:"required"`   // 目标操作系统
	Goarch  string   `json:"goarch" binding:"required"` // 目标架构
	Flags   []string `json:"flags"`                     // 编译标志
	Ldflags string   `json:"ldflags"`                   // 链接标志
	Tags    string   `json:"tags"`                      // 编译标签
	Scripts []string `json:"scripts"`                   // 编译前执行的JavaScript脚本，可以通过compileInfo访问编译配置
	EnvVars []string `json:"envVars"`                   // 环境变量
}

// NewCompileRequest 创建编译请求
func NewCompileRequest() *CompileRequest {
	return &CompileRequest{}
}

// Validate 验证参数
func (r *CompileRequest) Validate() error {
	return nil
}
