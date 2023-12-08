package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Controller interface {
	Update(ctx *Context) error
	Draw(screen *ebiten.Image) error
}
