package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Controller interface {
	Update(ctx *Context) error
	Draw(screen *ebiten.Image) error

	// Вызывается в самом начале
	OnAppInit(ctx *Context) error
}
