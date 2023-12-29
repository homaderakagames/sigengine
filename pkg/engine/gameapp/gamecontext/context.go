package gamecontext

type Context struct {
	storage map[any]any

	LogicalWidth  int
	LogicalHeight int
}

func NewContext() *Context {
	return &Context{storage: map[any]any{}}
}

func (c *Context) Put(k, v any) {
	c.storage[k] = v
}

func (c *Context) Get(k any) any {
	return c.storage[k]
}

func (c *Context) Unset(k any) {
	delete(c.storage, k)
}
