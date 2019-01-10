package main

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/popeyeio/gintool"
	"github.com/popeyeio/handy"
)

func main() {
	r := gin.New()
	r.RedirectTrailingSlash = false // turn off 301

	r.Use(gintool.SetRequestID(GetUUID))

	e := gintool.NewEngine()
	e.Use(CheckUser)

	r.POST("/records", e.GinHandler(CreateRecord))

	r.Run(":1990")
}

func GetUUID(c *gin.Context) string {
	u, _ := handy.GetUUID()
	return u
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
	gintool.MustBind(c, record, gintool.BHeader|gintool.BFormBody|gintool.BValidator)

	fmt.Println("deal with record")

	gc.Finish(gintool.CodeOKZero, nil)
}
