package binder

import (
	"github.com/gin-gonic/gin"
)

var FormMultipartBinder = &formMultipartBinder{}

type formMultipartBinder struct {
}

var _ Binder = (*formMultipartBinder)(nil)

func (formMultipartBinder) Name() string {
	return BNameFormMultipart
}

func (formMultipartBinder) Bind(c *gin.Context, obj interface{}) error {
	if err := c.Request.ParseMultipartForm(memoryMax); err != nil {
		return err
	}
	return bind(obj, c.Request.MultipartForm.Value, BNameForm, false)
}
