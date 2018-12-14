package gintool

import (
	"fmt"
	"sync"
)

type Kind string

var (
	KindHeader   = Kind("header")
	KindParam    = Kind("param")
	KindQuery    = Kind("query")
	KindPostForm = Kind("postform")
)

type EmptyError struct {
	kind Kind
	key  string
}

var _ error = (*EmptyError)(nil)

func NewEmptyError(kind Kind, key string) *EmptyError {
	return &EmptyError{
		kind: kind,
		key:  key,
	}
}

func NewHeaderEmptyError(key string) *EmptyError {
	return NewEmptyError(KindHeader, key)
}

func NewParamEmptyError(key string) *EmptyError {
	return NewEmptyError(KindParam, key)
}

func NewQueryEmptyError(key string) *EmptyError {
	return NewEmptyError(KindQuery, key)
}

func NewPostFormEmptyError(key string) *EmptyError {
	return NewEmptyError(KindPostForm, key)
}

func (e *EmptyError) Error() string {
	return fmt.Sprintf("%s(%s) is empty", e.key, e.kind)
}

type GSError struct {
	code int
	err  error
}

var _ error = (*GSError)(nil)

func (e *GSError) GetCode() int {
	return e.code
}

func (e *GSError) GetError() error {
	return e.err
}

func (e *GSError) Error() string {
	return fmt.Sprintf("%s - %+v", CodeMsg(e.code), e.err)
}

func (e *GSError) reset() {
	e.code = 0
	e.err = nil
}

var gsErrorPool sync.Pool

func AcquireGSError(code int, err error) (e *GSError) {
	v := gsErrorPool.Get()
	if v == nil {
		e = &GSError{}
	} else {
		e = v.(*GSError)
	}

	e.code = code
	e.err = err
	return
}

func ReleaseGSError(e *GSError) {
	if e != nil {
		e.reset()
		gsErrorPool.Put(e)
	}
}
