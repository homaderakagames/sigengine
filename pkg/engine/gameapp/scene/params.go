package scene

import (
	"github.com/homaderakagames/sigengine/pkg/engine/gameapp/gamecontext"
)

const ChangeSceneKey = "change_scene"

type Params struct {
	Scene Scene
}

func (p *Params) Get(ctx *gamecontext.Context) (ok bool) {
	if v := ctx.Get(ChangeSceneKey); v != nil {
		if casted, ok := v.(Scene); ok {
			p.Scene = casted
			return ok
		}
	}
	return false
}

func (p *Params) Set(ctx *gamecontext.Context) {
	ctx.Put(ChangeSceneKey, p.Scene)
}
