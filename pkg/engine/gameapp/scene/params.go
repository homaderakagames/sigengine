package scene

import "github.com/homaderakagames/sigengine/pkg/engine/gameapp/gamecontext"

const _changeSceneKey = "change_scene"

type Params struct {
	Scene Scene
}

func (p *Params) Get(ctx *gamecontext.Context) (ok bool) {
	if v := ctx.Value(_changeSceneKey); v != nil {
		if casted, ok := v.(Scene); ok {
			p.Scene = casted
			return ok
		}
	}
	return false
}

func (p *Params) Set(ctx *gamecontext.Context) {
	ctx.Put(_changeSceneKey, p.Scene)
}
