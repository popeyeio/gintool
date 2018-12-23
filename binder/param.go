package binder

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

const tagKeyParam = "param"

var ParamBinder = &paramBinder{}

type paramBinder struct {
}

var _ Binder = (*paramBinder)(nil)

func (paramBinder) Name() string {
	return tagKeyParam
}

func (paramBinder) Bind(c *gin.Context, obj interface{}) error {
	return bind(obj, parse(c.Params), tagKeyParam, false)
}

func parse(ps gin.Params) (v url.Values) {
	v = make(url.Values)
	for _, p := range ps {
		if _, exists := v[p.Key]; exists {
			v[p.Key] = append(v[p.Key], p.Value)
		} else {
			v[p.Key] = []string{p.Value}
		}
	}
	return
}
