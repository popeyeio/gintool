package gintool

import (
	"reflect"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

const (
	tagKeyValidator = "valid"
)

type gintoolValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = (*gintoolValidator)(nil)

func (v *gintoolValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyInit()
		return v.validate.Struct(obj)
	}
	return nil
}

func (v *gintoolValidator) RegisterValidation(key string, fn validator.Func) error {
	v.lazyInit()
	return v.validate.RegisterValidation(key, fn)
}

func (v *gintoolValidator) lazyInit() {
	v.once.Do(func() {
		config := &validator.Config{
			TagName: tagKeyValidator,
		}
		v.validate = validator.New(config)
	})
}

func kindOfData(data interface{}) reflect.Kind {
	rv := reflect.ValueOf(data)
	kind := rv.Kind()
	if kind == reflect.Ptr {
		kind = rv.Elem().Kind()
	}
	return kind
}

func CloseGinValidator() {
	binding.Validator = nil
}

var (
	GintoolValidator = &gintoolValidator{}
)

func Validate(v interface{}) error {
	return GintoolValidator.ValidateStruct(v)
}

func MustValidate(v interface{}) {
	MustDo(func() (interface{}, error) {
		return nil, Validate(v)
	}, CodeValidateErr)
}
