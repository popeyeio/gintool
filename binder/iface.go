package binder

import (
	"errors"
	"net/textproto"

	"github.com/gin-gonic/gin"
)

var ErrInvalidType = errors.New("invalid type")

type Binder interface {
	Name() string
	Bind(*gin.Context, interface{}) error
}

func canonicalKey(key string, canonical bool) string {
	if !canonical {
		return key
	}
	return textproto.CanonicalMIMEHeaderKey(key)
}

func convertValue(value string) string {
	if value != "" {
		return value
	}
	return "0"
}
