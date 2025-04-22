package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	*gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func New(ctx *gin.Context) Controller {
	return Controller{
		Context: ctx,
	}
}

func (ctl Controller) Success() {
	ctl.JSON(200, &Response{
		Code: 200,
		Msg:  "success",
		Data: nil,
	})
}

// 成功并且返回数据
func (ctl Controller) SuccessData(data any) {
	ctl.JSON(200, &Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

// 失败
func (ctl Controller) Fail(msg string) {
	ctl.JSON(200, &Response{
		Code: 400,
		Msg:  msg,
		Data: nil,
	})
}

// 失败 返回指定code
func (ctl Controller) FailWithCode(code int, msg string) {
	ctl.JSON(200, &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// InternalServerError
func (ctl Controller) InternalServerError(msg string) {
	ctl.FailWithCode(http.StatusInternalServerError, msg)
}

func (ctl Controller) FailWithError(err error) {
	if strings.Contains(err.Error(), "Token is expired") {
		ctl.Unauthorized("token 过期")
		return
	}

	ctl.Fail(err.Error())
}

func (ctl Controller) BadRequest(msg string) {
	ctl.FailWithCode(http.StatusBadRequest, msg)
}

// 401
func (ctl Controller) Unauthorized(msg string) {
	ctl.FailWithCode(http.StatusUnauthorized, msg)
}

// 403
func (ctl Controller) Forbidden(msg string) {
	ctl.FailWithCode(http.StatusForbidden, msg)
}
