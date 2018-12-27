package gintool

import (
	"github.com/gin-gonic/gin"
)

const (
	KeyRequestID = "gintool_request_id"
)

func SetRequestID(f func(*gin.Context) string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(KeyRequestID, f(c))
	}
}

func GetRequestID(c *gin.Context) string {
	return c.GetString(KeyRequestID)
}
