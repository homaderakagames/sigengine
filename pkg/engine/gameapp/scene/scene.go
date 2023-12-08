package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/homaderakagames/sigengine/pkg/engine/gameapp/gamecontext"
)

type Scene interface {
	Update(ctx *gamecontext.Context) error
	Draw(screen *ebiten.Image) error

	// OnLoad вызывается сразу после смены сцены
	OnLoad(ctx *gamecontext.Context) error

	// OnDestroy вызывается сразу перед сменой сцены
	OnDestroy(ctx *gamecontext.Context) error
}
