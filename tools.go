package gintool

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/go-querystring/query"
	"github.com/popeyeio/gintool/binder"
	"github.com/popeyeio/gintool/json"
	"github.com/popeyeio/gintool/validator"
	"github.com/popeyeio/handy"
)

type RunFunc func() (interface{}, error)
type CallbackFunc func(error)

func HeaderBool(c *gin.Context, key string) (bool, error) {
	if v := c.GetHeader(key); v == "" {
		return false, NewHeaderEmptyError(key)
	} else {
		return strconv.ParseBool(v)
	}
}

func MustHeaderBool(c *gin.Context, key string, cbs ...CallbackFunc) bool {
	result := MustDoCallback(func() (interface{}, error) {
		return HeaderBool(c, key)
	}, CodeBadRequest, cbs...)
	return result.(bool)
}

func HeaderInt64(c *gin.Context, key string) (int64, error) {
	if v := c.GetHeader(key); v == "" {
		return 0, NewHeaderEmptyError(key)
	} else {
		return strconv.ParseInt(v, 10, 64)
	}
}

func MustHeaderInt64(c *gin.Context, key string, cbs ...CallbackFunc) int64 {
	result := MustDoCallback(func() (interface{}, error) {
		return HeaderInt64(c, key)
	}, CodeBadRequest, cbs...)
	return result.(int64)
}

func HeaderUint64(c *gin.Context, key string) (uint64, error) {
	if v := c.GetHeader(key); v == "" {
		return 0, NewHeaderEmptyError(key)
	} else {
		return strconv.ParseUint(v, 10, 64)
	}
}

func MustHeaderUInt64(c *gin.Context, key string, cbs ...CallbackFunc) uint64 {
	result := MustDoCallback(func() (interface{}, error) {
		return HeaderUint64(c, key)
	}, CodeBadRequest, cbs...)
	return result.(uint64)
}

func HeaderString(c *gin.Context, key string) (string, error) {
	if v := c.GetHeader(key); v == "" {
		return "", NewHeaderEmptyError(key)
	} else {
		return v, nil
	}
}

func MustHeaderString(c *gin.Context, key string, cbs ...CallbackFunc) string {
	result := MustDoCallback(func() (interface{}, error) {
		return HeaderString(c, key)
	}, CodeBadRequest, cbs...)
	return result.(string)
}

func ParamBool(c *gin.Context, key string) (bool, error) {
	if v := c.Param(key); v == "" {
		return false, NewParamEmptyError(key)
	} else {
		return strconv.ParseBool(v)
	}
}

func MustParamBool(c *gin.Context, key string, cbs ...CallbackFunc) bool {
	result := MustDoCallback(func() (interface{}, error) {
		return ParamBool(c, key)
	}, CodeBadRequest, cbs...)
	return result.(bool)
}

func ParamInt64(c *gin.Context, key string) (int64, error) {
	if v := c.Param(key); v == "" {
		return 0, NewParamEmptyError(key)
	} else {
		return strconv.ParseInt(v, 10, 64)
	}
}

func MustParamInt64(c *gin.Context, key string, cbs ...CallbackFunc) int64 {
	result := MustDoCallback(func() (interface{}, error) {
		return ParamInt64(c, key)
	}, CodeBadRequest, cbs...)
	return result.(int64)
}

func ParamUint64(c *gin.Context, key string) (uint64, error) {
	if v := c.Param(key); v == "" {
		return 0, NewParamEmptyError(key)
	} else {
		return strconv.ParseUint(v, 10, 64)
	}
}

func MustParamUInt64(c *gin.Context, key string, cbs ...CallbackFunc) uint64 {
	result := MustDoCallback(func() (interface{}, error) {
		return ParamUint64(c, key)
	}, CodeBadRequest, cbs...)
	return result.(uint64)
}

func ParamString(c *gin.Context, key string) (string, error) {
	if v := c.Param(key); v == "" {
		return "", NewParamEmptyError(key)
	} else {
		return v, nil
	}
}

func MustParamString(c *gin.Context, key string, cbs ...CallbackFunc) string {
	result := MustDoCallback(func() (interface{}, error) {
		return ParamString(c, key)
	}, CodeBadRequest, cbs...)
	return result.(string)
}

func QueryBool(c *gin.Context, key string) (bool, error) {
	if v := c.Query(key); v == "" {
		return false, NewQueryEmptyError(key)
	} else {
		return strconv.ParseBool(v)
	}
}

func MustQueryBool(c *gin.Context, key string, cbs ...CallbackFunc) bool {
	result := MustDoCallback(func() (interface{}, error) {
		return QueryBool(c, key)
	}, CodeBadRequest, cbs...)
	return result.(bool)
}

func QueryInt64(c *gin.Context, key string) (int64, error) {
	if v := c.Query(key); v == "" {
		return 0, NewQueryEmptyError(key)
	} else {
		return strconv.ParseInt(v, 10, 64)
	}
}

func MustQueryInt64(c *gin.Context, key string, cbs ...CallbackFunc) int64 {
	result := MustDoCallback(func() (interface{}, error) {
		return QueryInt64(c, key)
	}, CodeBadRequest, cbs...)
	return result.(int64)
}

func QueryUint64(c *gin.Context, key string) (uint64, error) {
	if v := c.Query(key); v == "" {
		return 0, NewQueryEmptyError(key)
	} else {
		return strconv.ParseUint(v, 10, 64)
	}
}

func MustQueryUInt64(c *gin.Context, key string, cbs ...CallbackFunc) uint64 {
	result := MustDoCallback(func() (interface{}, error) {
		return QueryUint64(c, key)
	}, CodeBadRequest, cbs...)
	return result.(uint64)
}

func QueryString(c *gin.Context, key string) (string, error) {
	if v := c.Query(key); v == "" {
		return "", NewQueryEmptyError(key)
	} else {
		return v, nil
	}
}

func MustQueryString(c *gin.Context, key string, cbs ...CallbackFunc) string {
	result := MustDoCallback(func() (interface{}, error) {
		return QueryString(c, key)
	}, CodeBadRequest, cbs...)
	return result.(string)
}

func PostFormBool(c *gin.Context, key string) (bool, error) {
	if v := c.PostForm(key); v == "" {
		return false, NewPostFormEmptyError(key)
	} else {
		return strconv.ParseBool(v)
	}
}

func MustPostFormBool(c *gin.Context, key string, cbs ...CallbackFunc) bool {
	result := MustDoCallback(func() (interface{}, error) {
		return PostFormBool(c, key)
	}, CodeBadRequest, cbs...)
	return result.(bool)
}

func PostFormInt64(c *gin.Context, key string) (int64, error) {
	if v := c.PostForm(key); v == "" {
		return 0, NewPostFormEmptyError(key)
	} else {
		return strconv.ParseInt(v, 10, 64)
	}
}

func MustPostFormInt64(c *gin.Context, key string, cbs ...CallbackFunc) int64 {
	result := MustDoCallback(func() (interface{}, error) {
		return PostFormInt64(c, key)
	}, CodeBadRequest, cbs...)
	return result.(int64)
}

func PostFormUint64(c *gin.Context, key string) (uint64, error) {
	if v := c.PostForm(key); v == "" {
		return 0, NewPostFormEmptyError(key)
	} else {
		return strconv.ParseUint(v, 10, 64)
	}
}

func MustPostFormUInt64(c *gin.Context, key string, cbs ...CallbackFunc) uint64 {
	result := MustDoCallback(func() (interface{}, error) {
		return PostFormUint64(c, key)
	}, CodeBadRequest, cbs...)
	return result.(uint64)
}

func PostFormString(c *gin.Context, key string) (string, error) {
	if v := c.PostForm(key); v == "" {
		return "", NewPostFormEmptyError(key)
	} else {
		return v, nil
	}
}

func MustPostFormString(c *gin.Context, key string, cbs ...CallbackFunc) string {
	result := MustDoCallback(func() (interface{}, error) {
		return PostFormString(c, key)
	}, CodeBadRequest, cbs...)
	return result.(string)
}

func CloseGinValidator() {
	binding.Validator = nil
}

// BindValidator needs tag "valid" in fields of v.
func Validate(v interface{}) error {
	return validator.GintoolValidator.ValidateStruct(v)
}

func MustValidate(v interface{}, cbs ...CallbackFunc) {
	MustDoCallback(func() (interface{}, error) {
		return nil, Validate(v)
	}, CodeValidateErr, cbs...)
}

// BindHeader needs tag "header" in fields of v.
// The value of tag "header" is automatically converted to the canonical format.
func BindHeader(c *gin.Context, v interface{}) error {
	return binder.HeaderBinder.Bind(c, v)
}

func MustBindHeader(c *gin.Context, v interface{}, cbs ...CallbackFunc) {
	MustDoCallback(func() (interface{}, error) {
		return nil, BindHeader(c, v)
	}, CodeBindErr, cbs...)
}

// BindParam needs tag "param" in fields of v.
func BindParam(c *gin.Context, v interface{}) error {
	return binder.ParamBinder.Bind(c, v)
}

func MustBindParam(c *gin.Context, v interface{}, cbs ...CallbackFunc) {
	MustDoCallback(func() (interface{}, error) {
		return nil, BindParam(c, v)
	}, CodeBindErr, cbs...)
}

func JSONUseNumber(enabled bool) {
	binding.EnableDecoderUseNumber = enabled
}

func JSONBindBody(c *gin.Context, v interface{}) error {
	return c.ShouldBindWith(v, binding.JSON)
}

func MustJSONBindBody(c *gin.Context, v interface{}, cbs ...CallbackFunc) {
	MustDoCallback(func() (interface{}, error) {
		return nil, JSONBindBody(c, v)
	}, CodeBindErr, cbs...)
}

func XMLBindBody(c *gin.Context, v interface{}) error {
	return c.ShouldBindWith(v, binding.XML)
}

func MustXMLBindBody(c *gin.Context, v interface{}, cbs ...CallbackFunc) {
	MustDoCallback(func() (interface{}, error) {
		return nil, XMLBindBody(c, v)
	}, CodeBindErr, cbs...)
}

// FormBindQuery needs tag "form" in fields of v.
// NOTE: form does not support "-".
func FormBindQuery(c *gin.Context, v interface{}) error {
	return c.ShouldBindWith(v, binding.Query)
}

func MustFormBindQuery(c *gin.Context, v interface{}, cbs ...CallbackFunc) {
	MustDoCallback(func() (interface{}, error) {
		return nil, FormBindQuery(c, v)
	}, CodeBindErr, cbs...)
}

// FormBindBody needs tag "form" in fields of v.
// NOTE: form does not support "-".
func FormBindBody(c *gin.Context, v interface{}) error {
	return c.ShouldBindWith(v, binding.FormPost)
}

func MustFormBindBody(c *gin.Context, v interface{}, cbs ...CallbackFunc) {
	MustDoCallback(func() (interface{}, error) {
		return nil, FormBindBody(c, v)
	}, CodeBindErr, cbs...)
}

// FormBindQueryBody needs tag "form" in fields of v.
// NOTE: form does not support "-".
func FormBindQueryBody(c *gin.Context, v interface{}) error {
	return c.ShouldBindWith(v, binding.Form)
}

func MustFormBindQueryBody(c *gin.Context, v interface{}, cbs ...CallbackFunc) {
	MustDoCallback(func() (interface{}, error) {
		return nil, FormBindQueryBody(c, v)
	}, CodeBindErr, cbs...)
}

func PBBindBody(c *gin.Context, v interface{}) error {
	return c.ShouldBindWith(v, binding.ProtoBuf)
}

func MustPBBindBody(c *gin.Context, v interface{}, cbs ...CallbackFunc) {
	MustDoCallback(func() (interface{}, error) {
		return nil, PBBindBody(c, v)
	}, CodeBindErr, cbs...)
}

func MsgpackBindBody(c *gin.Context, v interface{}) error {
	return c.ShouldBindWith(v, binding.MsgPack)
}

func MustMsgpackBindBody(c *gin.Context, v interface{}, cbs ...CallbackFunc) {
	MustDoCallback(func() (interface{}, error) {
		return nil, MsgpackBindBody(c, v)
	}, CodeBindErr, cbs...)
}

// EncodeValues needs tag "url" in fields of v.
func EncodeValues(v interface{}) (url.Values, error) {
	return query.Values(v)
}

func MustEncodeValues(v interface{}, cbs ...CallbackFunc) url.Values {
	result := MustDoCallback(func() (interface{}, error) {
		return EncodeValues(v)
	}, CodeEncodeErr, cbs...)
	return result.(url.Values)
}

// EncodeJSON needs tag "json" in fields of v.
// Note: ReleaseBuffer needs to be called after EncodeJSON is called successfully.
func EncodeJSON(v interface{}) (*bytes.Buffer, error) {
	buffer := AcquireBuffer()
	if err := json.NewEncoder(buffer).Encode(v); err != nil {
		ReleaseBuffer(buffer)
		return nil, err
	}
	return buffer, nil
}

// MustEncodeJSON needs tag "json" in fields of v.
// Note: ReleaseBuffer needs to be called after MustEncodeJSON is called successfully.
func MustEncodeJSON(v interface{}, cbs ...CallbackFunc) *bytes.Buffer {
	result := MustDoCallback(func() (interface{}, error) {
		return EncodeJSON(v)
	}, CodeEncodeErr, cbs...)
	return result.(*bytes.Buffer)
}

func MustDo(run RunFunc, codes ...int) interface{} {
	code := CodeDownstreamErr
	if len(codes) > 0 {
		code = codes[0]
	}

	return MustDoCallback(run, code)
}

// MustDoCallback will call cbs if and only if run has error.
func MustDoCallback(run RunFunc, code int, cbs ...CallbackFunc) interface{} {
	if run == nil {
		return nil
	}

	result, err := run()
	if err == nil {
		return result
	}

	for _, cb := range cbs {
		cb(err)
	}

	if !IsGintoolError(err) {
		err = AcquireGintoolError(code, err)
	}
	panic(err)
}

const (
	BValidator = 1 << iota
	BHeader
	BParam
	BJSONBody
	BXMLBody
	BFormQuery
	BFormBody
	BFormQueryBody
	BPBBody
	BMsgpackBody
)

var funcs = map[int]func(*gin.Context, interface{}) error{
	BHeader:        BindHeader,
	BParam:         BindParam,
	BJSONBody:      JSONBindBody,
	BXMLBody:       XMLBindBody,
	BFormQuery:     FormBindQuery,
	BFormBody:      FormBindBody,
	BFormQueryBody: FormBindQueryBody,
	BPBBody:        PBBindBody,
	BMsgpackBody:   MsgpackBindBody,
}

func Bind(c *gin.Context, v interface{}, flag int) (err error) {
	for k, f := range funcs {
		if flag&k != 0 {
			if err = f(c, v); err != nil {
				return
			}
		}
	}

	// this must be the last one.
	if flag&BValidator != 0 {
		if err = Validate(v); err != nil {
			return AcquireGintoolError(CodeValidateErr, err)
		}
	}
	return
}

func MustBind(c *gin.Context, v interface{}, flag int, cbs ...CallbackFunc) {
	MustDoCallback(func() (interface{}, error) {
		return nil, Bind(c, v, flag)
	}, CodeBindErr, cbs...)
}

func GetRequestHost(req *http.Request) (host string) {
	if host = req.Host; handy.IsEmptyStr(host) {
		host = req.URL.Host
	}
	return
}
