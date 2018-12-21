package gintool

import (
	"net/http"
)

const (
	CodeOKZero         = 0
	CodeOK             = 2000
	CodeCreated        = 2001
	CodePartialContent = 2006

	CodeMultipleChoices   = 3000
	CodeMovedPermanently  = 3001
	CodeFound             = 3002
	CodeSeeOther          = 3003
	CodeNotModified       = 3004
	CodeUseProxy          = 3005
	CodeTemporaryRedirect = 3007
	CodePermanentRedirect = 3008

	CodeBadRequest      = 4000
	CodeUnauthorized    = 4001
	CodeForbidden       = 4003
	CodeNotFound        = 4004
	CodeTooManyRequests = 4029
	CodeValidateErr     = 4050

	CodeInternalErr        = 5000
	CodeServiceUnavailable = 5003
	CodeParseDataErr       = 5013
	CodeDBErr              = 5014
	CodeHbaseErr           = 5015
	CodeRedisErr           = 5016
	CodeDiskErr            = 5017
	CodeNSQErr             = 5030
	CodeKafkaErr           = 5031
	CodeBindErr            = 5040
	CodeEncodeErr          = 5041
	CodeDownstreamErr      = 5100
)

var codeMsg = map[int]string{
	CodeOKZero:         "success",
	CodeOK:             "success",
	CodeCreated:        "created",
	CodePartialContent: "partial content",

	CodeMultipleChoices:   "multiple choices",
	CodeMovedPermanently:  "moved permanently",
	CodeFound:             "found",
	CodeSeeOther:          "see other",
	CodeNotModified:       "not modified",
	CodeUseProxy:          "use proxy",
	CodeTemporaryRedirect: "temporary redirect",
	CodePermanentRedirect: "permanent redirect",

	CodeBadRequest:      "bad request",
	CodeUnauthorized:    "unauthorized",
	CodeForbidden:       "forbidden",
	CodeNotFound:        "not found",
	CodeTooManyRequests: "too many requests",
	CodeValidateErr:     "validate error",

	CodeInternalErr:        "internal error",
	CodeServiceUnavailable: "service unavailable",
	CodeParseDataErr:       "parse request data error",
	CodeDBErr:              "db error",
	CodeHbaseErr:           "hbase error",
	CodeRedisErr:           "redis error",
	CodeDiskErr:            "disk error",
	CodeNSQErr:             "nsq error",
	CodeKafkaErr:           "kafka error",
	CodeBindErr:            "bind error",
	CodeEncodeErr:          "encode error",
	CodeDownstreamErr:      "downstream error",
}

func CodeMsg(code int) string {
	if msg, exists := codeMsg[code]; exists {
		return msg
	}
	return "unknown code"
}

var httpStatus = map[int]int{
	CodeOKZero:         http.StatusOK,
	CodeOK:             http.StatusOK,
	CodeCreated:        http.StatusCreated,
	CodePartialContent: http.StatusPartialContent,

	CodeMultipleChoices:   http.StatusMultipleChoices,
	CodeMovedPermanently:  http.StatusMovedPermanently,
	CodeFound:             http.StatusFound,
	CodeSeeOther:          http.StatusSeeOther,
	CodeNotModified:       http.StatusNotModified,
	CodeUseProxy:          http.StatusUseProxy,
	CodeTemporaryRedirect: http.StatusTemporaryRedirect,
	CodePermanentRedirect: http.StatusPermanentRedirect,

	CodeBadRequest:      http.StatusBadRequest,
	CodeUnauthorized:    http.StatusUnauthorized,
	CodeForbidden:       http.StatusForbidden,
	CodeNotFound:        http.StatusNotFound,
	CodeTooManyRequests: http.StatusTooManyRequests,
	CodeValidateErr:     http.StatusBadRequest,

	CodeInternalErr:        http.StatusInternalServerError,
	CodeServiceUnavailable: http.StatusInternalServerError,
	CodeParseDataErr:       http.StatusInternalServerError,
	CodeDBErr:              http.StatusInternalServerError,
	CodeHbaseErr:           http.StatusInternalServerError,
	CodeRedisErr:           http.StatusInternalServerError,
	CodeDiskErr:            http.StatusInternalServerError,
	CodeNSQErr:             http.StatusInternalServerError,
	CodeKafkaErr:           http.StatusInternalServerError,
	CodeBindErr:            http.StatusInternalServerError,
	CodeEncodeErr:          http.StatusInternalServerError,
	CodeDownstreamErr:      http.StatusInternalServerError,
}

const (
	UnknownStatus = 900
)

func HTTPStatus(code int) int {
	if status, exists := httpStatus[code]; exists {
		return status
	}
	return UnknownStatus
}

func RegisterCode(code int, msg string, status int) bool {
	if _, exists := codeMsg[code]; exists {
		return false
	}
	if _, exists := httpStatus[code]; exists {
		return false
	}

	codeMsg[code] = msg
	httpStatus[code] = status
	return true
}
