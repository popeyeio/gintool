package gintool

import (
	"bytes"
	"net/url"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/go-querystring/query"
	"github.com/popeyeio/gintool/json"
)

func HeaderBool(c *gin.Context, key string) (bool, error) {
	if v := c.GetHeader(key); v == "" {
		return false, NewHeaderEmptyError(key)
	} else {
		return strconv.ParseBool(v)
	}
}

func MustHeaderBool(c *gin.Context, key string) bool {
	result := MustDo(func() (interface{}, error) {
		return HeaderBool(c, key)
	}, CodeBadRequest)
	return result.(bool)
}

func HeaderInt64(c *gin.Context, key string) (int64, error) {
	if v := c.GetHeader(key); v == "" {
		return 0, NewHeaderEmptyError(key)
	} else {
		return strconv.ParseInt(v, 10, 64)
	}
}

func MustHeaderInt64(c *gin.Context, key string) int64 {
	result := MustDo(func() (interface{}, error) {
		return HeaderInt64(c, key)
	}, CodeBadRequest)
	return result.(int64)
}

func HeaderUInt64(c *gin.Context, key string) (uint64, error) {
	if v := c.GetHeader(key); v == "" {
		return 0, NewHeaderEmptyError(key)
	} else {
		return strconv.ParseUint(v, 10, 64)
	}
}

func MustHeaderUInt64(c *gin.Context, key string) uint64 {
	result := MustDo(func() (interface{}, error) {
		return HeaderUInt64(c, key)
	}, CodeBadRequest)
	return result.(uint64)
}

func HeaderString(c *gin.Context, key string) (string, error) {
	if v := c.GetHeader(key); v == "" {
		return "", NewHeaderEmptyError(key)
	} else {
		return v, nil
	}
}

func MustHeaderString(c *gin.Context, key string) string {
	result := MustDo(func() (interface{}, error) {
		return HeaderString(c, key)
	}, CodeBadRequest)
	return result.(string)
}

func ParamBool(c *gin.Context, key string) (bool, error) {
	if v := c.Param(key); v == "" {
		return false, NewParamEmptyError(key)
	} else {
		return strconv.ParseBool(v)
	}
}

func MustParamBool(c *gin.Context, key string) bool {
	result := MustDo(func() (interface{}, error) {
		return ParamBool(c, key)
	}, CodeBadRequest)
	return result.(bool)
}

func ParamInt64(c *gin.Context, key string) (int64, error) {
	if v := c.Param(key); v == "" {
		return 0, NewParamEmptyError(key)
	} else {
		return strconv.ParseInt(v, 10, 64)
	}
}

func MustParamInt64(c *gin.Context, key string) int64 {
	result := MustDo(func() (interface{}, error) {
		return ParamInt64(c, key)
	}, CodeBadRequest)
	return result.(int64)
}

func ParamUInt64(c *gin.Context, key string) (uint64, error) {
	if v := c.Param(key); v == "" {
		return 0, NewParamEmptyError(key)
	} else {
		return strconv.ParseUint(v, 10, 64)
	}
}

func MustParamUInt64(c *gin.Context, key string) uint64 {
	result := MustDo(func() (interface{}, error) {
		return ParamUInt64(c, key)
	}, CodeBadRequest)
	return result.(uint64)
}

func ParamString(c *gin.Context, key string) (string, error) {
	if v := c.Param(key); v == "" {
		return "", NewParamEmptyError(key)
	} else {
		return v, nil
	}
}

func MustParamString(c *gin.Context, key string) string {
	result := MustDo(func() (interface{}, error) {
		return ParamString(c, key)
	}, CodeBadRequest)
	return result.(string)
}

func QueryBool(c *gin.Context, key string) (bool, error) {
	if v := c.Query(key); v == "" {
		return false, NewQueryEmptyError(key)
	} else {
		return strconv.ParseBool(v)
	}
}

func MustQueryBool(c *gin.Context, key string) bool {
	result := MustDo(func() (interface{}, error) {
		return QueryBool(c, key)
	}, CodeBadRequest)
	return result.(bool)
}

func QueryInt64(c *gin.Context, key string) (int64, error) {
	if v := c.Query(key); v == "" {
		return 0, NewQueryEmptyError(key)
	} else {
		return strconv.ParseInt(v, 10, 64)
	}
}

func MustQueryInt64(c *gin.Context, key string) int64 {
	result := MustDo(func() (interface{}, error) {
		return QueryInt64(c, key)
	}, CodeBadRequest)
	return result.(int64)
}

func QueryUInt64(c *gin.Context, key string) (uint64, error) {
	if v := c.Query(key); v == "" {
		return 0, NewQueryEmptyError(key)
	} else {
		return strconv.ParseUint(v, 10, 64)
	}
}

func MustQueryUInt64(c *gin.Context, key string) uint64 {
	result := MustDo(func() (interface{}, error) {
		return QueryUInt64(c, key)
	}, CodeBadRequest)
	return result.(uint64)
}

func QueryString(c *gin.Context, key string) (string, error) {
	if v := c.Query(key); v == "" {
		return "", NewQueryEmptyError(key)
	} else {
		return v, nil
	}
}

func MustQueryString(c *gin.Context, key string) string {
	result := MustDo(func() (interface{}, error) {
		return QueryString(c, key)
	}, CodeBadRequest)
	return result.(string)
}

func PostFormBool(c *gin.Context, key string) (bool, error) {
	if v := c.PostForm(key); v == "" {
		return false, NewPostFormEmptyError(key)
	} else {
		return strconv.ParseBool(v)
	}
}

func MustPostFormBool(c *gin.Context, key string) bool {
	result := MustDo(func() (interface{}, error) {
		return PostFormBool(c, key)
	}, CodeBadRequest)
	return result.(bool)
}

func PostFormInt64(c *gin.Context, key string) (int64, error) {
	if v := c.PostForm(key); v == "" {
		return 0, NewPostFormEmptyError(key)
	} else {
		return strconv.ParseInt(v, 10, 64)
	}
}

func MustPostFormInt64(c *gin.Context, key string) int64 {
	result := MustDo(func() (interface{}, error) {
		return PostFormInt64(c, key)
	}, CodeBadRequest)
	return result.(int64)
}

func PostFormUInt64(c *gin.Context, key string) (uint64, error) {
	if v := c.PostForm(key); v == "" {
		return 0, NewPostFormEmptyError(key)
	} else {
		return strconv.ParseUint(v, 10, 64)
	}
}

func MustPostFormUInt64(c *gin.Context, key string) uint64 {
	result := MustDo(func() (interface{}, error) {
		return PostFormUInt64(c, key)
	}, CodeBadRequest)
	return result.(uint64)
}

func PostFormString(c *gin.Context, key string) (string, error) {
	if v := c.PostForm(key); v == "" {
		return "", NewPostFormEmptyError(key)
	} else {
		return v, nil
	}
}

func MustPostFormString(c *gin.Context, key string) string {
	result := MustDo(func() (interface{}, error) {
		return PostFormString(c, key)
	}, CodeBadRequest)
	return result.(string)
}

// BindHeader needs tag "header" in fields of v.
// The value of tag "header" is automatically converted to the canonical format.
func BindHeader(c *gin.Context, v interface{}) error {
	return c.ShouldBindWith(v, Header)
}

func MustBindHeader(c *gin.Context, v interface{}) {
	MustDo(func() (interface{}, error) {
		return nil, BindHeader(c, v)
	}, CodeBindErr)
}

func JSONUseNumber(enabled bool) {
	binding.EnableDecoderUseNumber = enabled
}

func JSONBindBody(c *gin.Context, v interface{}) error {
	return c.ShouldBindWith(v, binding.JSON)
}

func MustJSONBindBody(c *gin.Context, v interface{}) {
	MustDo(func() (interface{}, error) {
		return nil, JSONBindBody(c, v)
	}, CodeBindErr)
}

func XMLBindBody(c *gin.Context, v interface{}) error {
	return c.ShouldBindWith(v, binding.XML)
}

func MustXMLBindBody(c *gin.Context, v interface{}) {
	MustDo(func() (interface{}, error) {
		return nil, XMLBindBody(c, v)
	}, CodeBindErr)
}

// FormBindQuery needs tag "form" in fields of v.
// NOTE: form does not support "-".
func FormBindQuery(c *gin.Context, v interface{}) error {
	return c.ShouldBindWith(v, binding.Query)
}

func MustFormBindQuery(c *gin.Context, v interface{}) {
	MustDo(func() (interface{}, error) {
		return nil, FormBindQuery(c, v)
	}, CodeBindErr)
}

// FormBindBody needs tag "form" in fields of v.
// NOTE: form does not support "-".
func FormBindBody(c *gin.Context, v interface{}) error {
	return c.ShouldBindWith(v, binding.FormPost)
}

func MustFormBindBody(c *gin.Context, v interface{}) {
	MustDo(func() (interface{}, error) {
		return nil, FormBindBody(c, v)
	}, CodeBindErr)
}

// FormBindQueryBody needs tag "form" in fields of v.
// NOTE: form does not support "-".
func FormBindQueryBody(c *gin.Context, v interface{}) error {
	return c.ShouldBindWith(v, binding.Form)
}

func MustFormBindQueryBody(c *gin.Context, v interface{}) {
	MustDo(func() (interface{}, error) {
		return nil, FormBindQueryBody(c, v)
	}, CodeBindErr)
}

func PBBindBody(c *gin.Context, v interface{}) error {
	return c.ShouldBindWith(v, binding.ProtoBuf)
}

func MustPBBindBody(c *gin.Context, v interface{}) {
	MustDo(func() (interface{}, error) {
		return nil, PBBindBody(c, v)
	}, CodeBindErr)
}

func MsgpackBindBody(c *gin.Context, v interface{}) error {
	return c.ShouldBindWith(v, binding.MsgPack)
}

func MustMsgpackBindBody(c *gin.Context, v interface{}) {
	MustDo(func() (interface{}, error) {
		return nil, MsgpackBindBody(c, v)
	}, CodeBindErr)
}

// EncodeValues needs tag "url" in fields of v.
func EncodeValues(v interface{}) (url.Values, error) {
	return query.Values(v)
}

func MustEncodeValues(v interface{}) url.Values {
	result := MustDo(func() (interface{}, error) {
		return EncodeValues(v)
	}, CodeEncodeErr)
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
// Note: ReleaseBuffer needs to be called after MustEncodeJSON.
func MustEncodeJSON(v interface{}) *bytes.Buffer {
	result := MustDo(func() (interface{}, error) {
		return EncodeJSON(v)
	}, CodeEncodeErr)
	return result.(*bytes.Buffer)
}

func MustDo(f func() (interface{}, error), codes ...int) interface{} {
	code := CodeDownstreamErr
	if len(codes) > 0 {
		code = codes[0]
	}

	result, err := f()
	if err != nil {
		panic(AcquireGintoolError(code, err))
	}
	return result
}

func Bytes2Str(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := &reflect.StringHeader{
		Data: bh.Data,
		Len:  bh.Len,
	}
	return *(*string)(unsafe.Pointer(sh))
}

func Str2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := &reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(bh))
}
