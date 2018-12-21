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
	c.JSON(gintool.HTTPStatus(code), gintool.RespOK(code, data))
}

func AbortWithCodeErr(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(gintool.HTTPStatus(code), gintool.RespError(code, err))
}

func CheckUser(c *gin.Context, gc *gintool.Context) {
	user := gintool.MustHeaderString(c, "X-User")
	if user != "popeye" {
		gc.Abort(gintool.CodeBadRequest, errors.New("invalid user"))
	}
}

type Record struct {
	AccountID int64   `header:"X-Accountid" form:"-" valid:"required"`
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
