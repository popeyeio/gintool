package binder

import (
	"github.com/gin-gonic/gin"
)

var QueryBinder = &queryBinder{}

type queryBinder struct {
}

var _ Binder = (*queryBinder)(nil)

func (queryBinder) Name() string {
	return BNameQuery
}

func (queryBinder) Bind(c *gin.Context, obj interface{}) error {
	return bind(obj, c.Request.URL.Query(), BNameForm, false)
}
