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
	*c = *ctx.(*Context)
	return c, cancel
}
