package gintool

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(*gin.Context, *Context)
type HandlerFuncsChain []HandlerFunc

type Engine struct {
	middlewares HandlerFuncsChain
	finisher    HandlerFunc
	aborter     HandlerFunc
	contextPool sync.Pool
}

type Option func(*Engine)

func WithFinisher(finisher HandlerFunc) Option {
	return func(e *Engine) {
		if finisher != nil {
			e.finisher = finisher
		}
	}
}

func WithAborter(aborter HandlerFunc) Option {
	return func(e *Engine) {
		if aborter != nil {
			e.aborter = aborter
		}
	}
}

func NewEngine(opts ...Option) *Engine {
	e := &Engine{
		finisher: func(*gin.Context, *Context) {},
		aborter:  func(*gin.Context, *Context) {},
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

func (e *Engine) Use(middleware ...HandlerFunc) {
	e.middlewares = append(e.middlewares, middleware...)
}

func (e *Engine) GinHandler(handlers ...HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		gc := e.acquireContext()
		defer e.releaseContext(gc)

		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(*GintoolError); ok {
					gc.Abort(err.GetCode(), err.GetError())
					ReleaseGintoolError(err)
					e.aborter(c, gc)
				} else {
					panic(r)
				}
			}
		}()

		gc.handlers = append(gc.handlers, e.middlewares...)
		gc.handlers = append(gc.handlers, handlers...)

		for _, handler := range gc.handlers {
			handler(c, gc)

			if !gc.IsOK() {
				break
			}
		}

		if gc.IsOK() {
			e.finisher(c, gc)
		} else {
			e.aborter(c, gc)
		}
	}
}

func (e *Engine) acquireContext() (gc *Context) {
	v := e.contextPool.Get()
	if v == nil {
		gc = &Context{}
	} else {
		gc = v.(*Context)
	}

	gc.engine = e
	gc.ok = true
	return
}

func (e *Engine) releaseContext(gc *Context) {
	if gc != nil {
		gc.reset()
		e.contextPool.Put(gc)
	}
}
