package v1

import (
	"fmt"
	"lime/internal/common/controller"
	"lime/internal/global"
	"lime/pkg/utils"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Upload doc
//
//	@tags			基础接口
//	@Summary		上传文件
//	@Description	上传文件
//	@Produce		json
//	@Param			data	body		requests.LoginRequest					true	"用户名、密码"
//	@Success		200		{object}	controller.Response{Data=response.Auth}	"code: 200 成功"
//	@Failure		500		{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/admin/upload [post]
func Upload(c *gin.Context) {
	ctl := controller.New(c)

	file, err := c.FormFile("file")
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	global.Logger.Sugar().Infof("上传文件: %s", file.Filename)

	// 去除文件的后缀名
	name := strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename))
	// 后缀名
	ext := filepath.Ext(file.Filename)
	// 文件名称 + 时间戳 再 hash
	name = fmt.Sprintf("%s_%d", name, time.Now().UnixNano())
	// 对文件进行重命名
	name = utils.ShaString(name)
	// 只要前32位
	name = name[:32]
	dst := filepath.Join(global.ROOT_PATH, "static", name+ext)
	// 上传文件至指定的完整文件路径
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(map[string]any{
		"url": fmt.Sprintf("/static/%s", name+ext),
	})
}
