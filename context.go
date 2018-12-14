package gintool

type Context struct {
	engine   *Engine
	handlers HandlerFuncsChain

	ok   bool
	code int
	data interface{}
	err  error
}

func (c *Context) Finish(code int, data interface{}) {
	c.ok = true
	c.code = code
	c.data = data
}

func (c *Context) Abort(code int, err error) {
	c.ok = false
	c.code = code
	c.err = err
}

func (c *Context) IsOK() bool {
	return c.ok
}

func (c *Context) GetCode() int {
	return c.code
}

func (c *Context) GetData() interface{} {
	return c.data
}

func (c *Context) GetError() error {
	return c.err
}

func (c *Context) reset() {
	c.engine = nil
	c.handlers = c.handlers[:0]
	c.ok = false
	c.code = 0
	c.data = nil
	c.err = nil
}
