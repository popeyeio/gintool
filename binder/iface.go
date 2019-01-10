package binder

import (
	"github.com/gin-gonic/gin"
)

type Binder interface {
	Name() string
	Bind(*gin.Context, interface{}) error
}
