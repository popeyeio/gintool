package gintool

import (
	"fmt"
	"sync"
)

type Kind int

const (
	KindHeader Kind = iota
	KindParam
	KindQuery
	KindPostForm
)

var _ fmt.Stringer = (*Kind)(nil)

func (k Kind) String() string {
	switch k {
	case KindHeader:
		return "header"
	case KindParam:
		return "param"
	case KindQuery:
		return "query"
	case KindPostForm:
		return "postform"
	}
	return fmt.Sprintf("unknown kind: %d", k)
}

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

func (e EmptyError) Error() string {
	return fmt.Sprintf("%s(%s) is empty", e.key, e.kind.String())
}

type GintoolError struct {
	code int
	err  error
}

var _ error = (*GintoolError)(nil)

func (e GintoolError) GetCode() int {
	return e.code
}

func (e GintoolError) GetError() error {
	return e.err
}

func (e GintoolError) Error() string {
	return fmt.Sprintf("%s - %+v", CodeMsg(e.code), e.err)
}

func (e *GintoolError) reset() {
	e.code = 0
	e.err = nil
}

var gintoolErrorPool sync.Pool

func AcquireGintoolError(code int, err error) (e *GintoolError) {
	if v := gintoolErrorPool.Get(); v != nil {
		e = v.(*GintoolError)
	} else {
		e = &GintoolError{}
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

func IsGintoolError(err error) bool {
	_, ok := err.(*GintoolError)
	return ok
}
