package binder

import (
	"errors"
	"net/textproto"

	"github.com/gin-gonic/gin"
	"github.com/popeyeio/handy"
)

const (
	memoryMax = 1 << 25

	BNameHeader        = "header"
	BNameParam         = "param"
	BNameQuery         = "query"
	BNameFormPost      = "form-urlencoded"
	BNameForm          = "form"
	BNameFormMultipart = "multipart/form-data"
)

var (
	ErrInvalidType = errors.New("invalid type")
)

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
	if !handy.IsEmptyStr(value) {
		return value
	}
	return "0"
}
