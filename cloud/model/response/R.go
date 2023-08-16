package R

import "github.com/gin-gonic/gin"

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func result(httpCode int, data interface{}, message string, ctx *gin.Context) {
	ctx.JSON(httpCode, Response{
		Data:    data,
		Message: message,
	})
}

// OK 操作成功，返回 200 状态码。提示：0K 状态码返回函数分两种，一种是返回 200 的, 一种是返回 201 的，当成功的在服务器中创建了资源，应当使用返回 201 状态码的函数。
func OK(data interface{}, message string, ctx *gin.Context) {
	if message != "" {
		result(200, data, message, ctx)
	} else {
		result(200, data, "操作成功", ctx)
	}
}

// CreateOk 资源创建成功，返回 201 状态码
func CreateOk(data interface{}, message string, ctx *gin.Context) {
	if message != "" {
		result(201, data, message, ctx)
	} else {
		result(201, data, "操作成功", ctx)
	}
}

// FailBadRequest 400 错误
func FailBadRequest(data interface{}, message string, ctx *gin.Context) {
	if message != "" {
		result(400, data, message, ctx)
	} else {
		result(400, data, "操作失败", ctx)
	}
}

// FailServerError 服务器处理异常
func FailServerError(data interface{}, message string, ctx *gin.Context) {
	if message != "" {
		result(500, data, message, ctx)
	} else {
		result(500, data, "服务器错误", ctx)
	}
}
