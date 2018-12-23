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
	"github.com/popeyeio/gintool/binder"
	"github.com/popeyeio/gintool/json"
	"github.com/popeyeio/gintool/validator"
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

func HeaderUint64(c *gin.Context, key string) (uint64, error) {
	if v := c.GetHeader(key); v == "" {
		return 0, NewHeaderEmptyError(key)
	} else {
		return strconv.ParseUint(v, 10, 64)
	}
}

func MustHeaderUInt64(c *gin.Context, key string) uint64 {
	result := MustDo(func() (interface{}, error) {
		return HeaderUint64(c, key)
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

func ParamUint64(c *gin.Context, key string) (uint64, error) {
	if v := c.Param(key); v == "" {
		return 0, NewParamEmptyError(key)
	} else {
		return strconv.ParseUint(v, 10, 64)
	}
}

func MustParamUInt64(c *gin.Context, key string) uint64 {
	result := MustDo(func() (interface{}, error) {
		return ParamUint64(c, key)
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

func QueryUint64(c *gin.Context, key string) (uint64, error) {
	if v := c.Query(key); v == "" {
		return 0, NewQueryEmptyError(key)
	} else {
		return strconv.ParseUint(v, 10, 64)
	}
}

func MustQueryUInt64(c *gin.Context, key string) uint64 {
	result := MustDo(func() (interface{}, error) {
		return QueryUint64(c, key)
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

func PostFormUint64(c *gin.Context, key string) (uint64, error) {
	if v := c.PostForm(key); v == "" {
		return 0, NewPostFormEmptyError(key)
	} else {
		return strconv.ParseUint(v, 10, 64)
	}
}

func MustPostFormUInt64(c *gin.Context, key string) uint64 {
	result := MustDo(func() (interface{}, error) {
		return PostFormUint64(c, key)
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

func CloseGinValidator() {
	binding.Validator = nil
}

// BindValidator needs tag "valid" in fields of v.
func Validate(v interface{}) error {
	return validator.GintoolValidator.ValidateStruct(v)
}

func MustValidate(v interface{}) {
	MustDo(func() (interface{}, error) {
		return nil, Validate(v)
	}, CodeValidateErr)
}

// BindHeader needs tag "header" in fields of v.
// The value of tag "header" is automatically converted to the canonical format.
func BindHeader(c *gin.Context, v interface{}) error {
	return binder.HeaderBinder.Bind(c, v)
}

func MustBindHeader(c *gin.Context, v interface{}) {
	MustDo(func() (interface{}, error) {
		return nil, BindHeader(c, v)
	}, CodeBindErr)
}

// BindParam needs tag "param" in fields of v.
func BindParam(c *gin.Context, v interface{}) error {
	return binder.ParamBinder.Bind(c, v)
}

func MustBindParam(c *gin.Context, v interface{}) {
	MustDo(func() (interface{}, error) {
		return nil, BindParam(c, v)
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
	if err == nil {
		return result
	}
	if IsGintoolError(err) {
		panic(err)
	}
	panic(AcquireGintoolError(code, err))
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

func Bind(c *gin.Context, v interface{}, flag int) (err error) {
	if flag&BHeader != 0 {
		if err = BindHeader(c, v); err != nil {
			return
		}
	}
	if flag&BParam != 0 {
		if err = BindParam(c, v); err != nil {
			return
		}
	}
	if flag&BJSONBody != 0 {
		if err = JSONBindBody(c, v); err != nil {
			return
		}
	}
	if flag&BXMLBody != 0 {
		if err = XMLBindBody(c, v); err != nil {
			return
		}
	}
	if flag&BFormQuery != 0 {
		if err = FormBindQuery(c, v); err != nil {
			return
		}
	}
	if flag&BFormBody != 0 {
		if err = FormBindBody(c, v); err != nil {
			return
		}
	}
	if flag&BFormQueryBody != 0 {
		if err = FormBindQueryBody(c, v); err != nil {
			return
		}
	}
	if flag&BPBBody != 0 {
		if err = PBBindBody(c, v); err != nil {
			return
		}
	}
	if flag&BMsgpackBody != 0 {
		if err = MsgpackBindBody(c, v); err != nil {
			return
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

func MustBind(c *gin.Context, v interface{}, flag int) {
	MustDo(func() (interface{}, error) {
		return nil, Bind(c, v, flag)
	}, CodeBindErr)
}
