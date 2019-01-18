package binder

import (
	"github.com/gin-gonic/gin"
)

var HeaderBinder = &headerBinder{}

type headerBinder struct {
}

var _ Binder = (*headerBinder)(nil)

func (headerBinder) Name() string {
	return BNameHeader
}

func (headerBinder) Bind(c *gin.Context, obj interface{}) error {
	return bind(obj, c.Request.Header, BNameHeader, true)
}
