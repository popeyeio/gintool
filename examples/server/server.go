package main

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/popeyeio/gintool"
)

func main() {
	e := NewGintoolEngine()
	e.Use(CheckUser)

	r := gin.New()

	r.POST("/records", e.GinHandler(CreateRecord))

	r.Run(":1990")
}

func NewGintoolEngine() *gintool.Engine {
	opts := []gintool.Option{
		gintool.WithFinisher(func(c *gin.Context, gc *gintool.Context) {
			FinishWithCodeData(c, gc.GetCode(), gc.GetData())
		}),
		gintool.WithAborter(func(c *gin.Context, gc *gintool.Context) {
			AbortWithCodeErr(c, gc.GetCode(), gc.GetError())
		}),
	}
	return gintool.NewEngine(opts...)
}

func FinishWithCodeData(c *gin.Context, code int, data interface{}) {
	c.JSON(gintool.HTTPStatus(code), RespOK(code, data))
}

func AbortWithCodeErr(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(gintool.HTTPStatus(code), RespError(code, err))
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RespOK(code int, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: gintool.CodeMsg(code),
		Data:    data,
	}
}

func RespError(code int, err error) *Response {
	resp := &Response{
		Code:    code,
		Message: gintool.CodeMsg(code),
	}
	if err != nil {
		resp.Message += " - " + err.Error()
	}
	return resp
}

func CheckUser(c *gin.Context, gc *gintool.Context) {
	user := gintool.MustHeaderString(c, "X-User")
	if user != "popeye" {
		gc.Abort(gintool.CodeBadRequest, errors.New("invalid user"))
	}
}

type Record struct {
	AccountID int64   `header:"X-AccountID" form:"-" valid:"required"`
	Action    string  `header:"-" form:"Action" valid:"required"`
	Amount    float32 `header:"-" form:"Amount" valid:"required"`
}

func CreateRecord(c *gin.Context, gc *gintool.Context) {
	record := &Record{}
	gintool.MustBindHeader(c, record)
	gintool.MustFormBindBody(c, record)
	gintool.MustValidate(record)

	fmt.Println("deal with record")

	gc.Finish(gintool.CodeOKZero, nil)
}
