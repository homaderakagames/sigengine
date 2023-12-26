package resources

import (
	"github.com/homaderakagames/sigengine/pkg/engine/gameapp/gamecontext"
)

const (
	GetResourceManager = "get_resource_manager"
)

type Params struct {
	ResourceManager *ResourceManager
}

func (p *Params) Get(ctx *gamecontext.Context) (ok bool) {
	if v := ctx.Get(GetResourceManager); v != nil {
		if casted, ok := v.(*ResourceManager); ok {
			p.ResourceManager = casted
			return ok
		}
	}

	return false
}

func (p *Params) Set(ctx *gamecontext.Context) {
	ctx.Put(GetResourceManager, p.ResourceManager)
}
