package gintool

import (
	"github.com/gin-gonic/gin"
)

type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RespOK(code int, data interface{}) *CommonResponse {
	return &CommonResponse{
		Code:    code,
		Message: CodeMsg(code),
		Data:    data,
	}
}

func RespError(code int, err error) *CommonResponse {
	resp := &CommonResponse{
		Code:    code,
		Message: CodeMsg(code),
	}
	if err != nil {
		resp.Message += " - " + err.Error()
	}
	return resp
}

func FinishWithCodeData(c *gin.Context, code int, data interface{}) {
	c.JSON(HTTPStatus(code), RespOK(code, data))
}

func AbortWithCodeErr(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(HTTPStatus(code), RespError(code, err))
}

func GetCommonFinisher() HandlerFunc {
	return func(c *gin.Context, gc *Context) {
		FinishWithCodeData(c, gc.GetCode(), gc.GetData())
	}
}

func GetCommonAborter() HandlerFunc {
	return func(c *gin.Context, gc *Context) {
		AbortWithCodeErr(c, gc.GetCode(), gc.GetError())
	}
}
