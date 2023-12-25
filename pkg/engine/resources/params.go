package resources

import (
	"github.com/homaderakagames/sigengine/pkg/engine/gameapp/gamecontext"
	fontManager "github.com/homaderakagames/sigengine/pkg/engine/resources/fonts"
)

const GetFontManagerKey = "get_font_manager"

type Params struct {
	FontManager *fontManager.FontManager
}

func (p *Params) Get(ctx *gamecontext.Context) (ok bool) {
	if v := ctx.Get(GetFontManagerKey); v != nil {
		if casted, ok := v.(*fontManager.FontManager); ok {
			p.FontManager = casted
			return ok
		}
	}
	return false
}

func (p *Params) Set(ctx *gamecontext.Context) {
	ctx.Put(GetFontManagerKey, p.FontManager)
}
