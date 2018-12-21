package main

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/popeyeio/gintool"
)

func main() {
	e := gintool.NewEngine()
	e.Use(CheckUser)

	r := gin.New()

	r.POST("/records", e.GinHandler(CreateRecord))

	r.Run(":1990")
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
