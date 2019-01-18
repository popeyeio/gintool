package binder

import (
	"github.com/gin-gonic/gin"
)

var FormBinder = &formBinder{}

type formBinder struct {
}

var _ Binder = (*formBinder)(nil)

func (formBinder) Name() string {
	return BNameForm
}

func (formBinder) Bind(c *gin.Context, obj interface{}) error {
	if err := c.Request.ParseForm(); err != nil {
		return err
	}
	c.Request.ParseMultipartForm(memoryMax)
	return bind(obj, c.Request.Form, BNameForm, false)
}
