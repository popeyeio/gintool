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

type GintoolError struct {
	code int
	err  error
}

var _ error = (*GintoolError)(nil)

func (e *GintoolError) GetCode() int {
	return e.code
}

func (e *GintoolError) GetError() error {
	return e.err
}

func (e *GintoolError) Error() string {
	return fmt.Sprintf("%s - %+v", CodeMsg(e.code), e.err)
}

func (e *GintoolError) reset() {
	e.code = 0
	e.err = nil
}

var gintoolErrorPool sync.Pool

func AcquireGintoolError(code int, err error) (e *GintoolError) {
	v := gintoolErrorPool.Get()
	if v == nil {
		e = &GintoolError{}
	} else {
		e = v.(*GintoolError)
	}

	e.code = code
	e.err = err
	return
}

func ReleaseGintoolError(e *GintoolError) {
	if e != nil {
		e.reset()
		gintoolErrorPool.Put(e)
	}
}
