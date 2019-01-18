package binder

import (
	"github.com/gin-gonic/gin"
)

var FormPostBinder = &formPostBinder{}

type formPostBinder struct {
}

var _ Binder = (*formPostBinder)(nil)

func (formPostBinder) Name() string {
	return BNameFormPost
}

func (formPostBinder) Bind(c *gin.Context, obj interface{}) error {
	if err := c.Request.ParseForm(); err != nil {
		return err
	}
	return bind(obj, c.Request.PostForm, BNameForm, false)
}
