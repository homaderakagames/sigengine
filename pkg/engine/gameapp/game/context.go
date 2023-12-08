package game

import "context"

type Context struct {
	context.Context

	LogicalWidth  int
	LogicalHeight int
}

func NewContext(ctx context.Context) *Context {
	return &Context{Context: ctx}
}

func (c *Context) WithCancel() (*Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(c)
	*c = *NewContext(ctx)
	return c, cancel
}

func (c *Context) Put(k, v any) {
	ctxVal := context.WithValue(c, k, v)
	*c = *NewContext(ctxVal)
}

func (c *Context) Get(k any) any {
	return c.Value(k)
}
