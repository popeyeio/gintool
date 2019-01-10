package binder

import (
	"github.com/gin-gonic/gin"
)

const TagKeyHeader = "header"

var HeaderBinder = &headerBinder{}

type headerBinder struct {
}

var _ Binder = (*headerBinder)(nil)

func (headerBinder) Name() string {
	return TagKeyHeader
}

func (headerBinder) Bind(c *gin.Context, obj interface{}) error {
	return bind(obj, c.Request.Header, TagKeyHeader, true)
}
