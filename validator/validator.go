package validator

import (
	"reflect"
	"sync"

	"gopkg.in/go-playground/validator.v8"
)

var GintoolValidator = &gintoolValidator{}

type gintoolValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ Validator = (*gintoolValidator)(nil)

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

func (v *gintoolValidator) RegisterAliasValidation(alias, tags string) {
	v.lazyInit()
	v.validate.RegisterAliasValidation(alias, tags)
}

func (v *gintoolValidator) RegisterStructValidation(fn validator.StructLevelFunc, types ...interface{}) {
	v.lazyInit()
	v.validate.RegisterStructValidation(fn, types...)
}

func (v *gintoolValidator) RegisterCustomTypeFunc(fn validator.CustomTypeFunc, types ...interface{}) {
	v.lazyInit()
	v.validate.RegisterCustomTypeFunc(fn, types...)
}

func (v *gintoolValidator) lazyInit() {
	v.once.Do(func() {
		config := &validator.Config{
			TagName: VTagValid,
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
