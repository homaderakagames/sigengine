package game

import (
	"context"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct {
	controller Controller

	context *Context
	cancel  context.CancelFunc
}

func NewGame(
	controller Controller,
	ctx *Context,
	cancel context.CancelFunc,
) *Game {

	return &Game{
		controller: controller,
		context:    ctx,
		cancel:     cancel,
	}
}

func (g *Game) Update() error {
	select {
	case <-g.context.Done():
		return ebiten.Termination
	default:
	}
	return g.controller.Update(g.context)
}

func (g *Game) Draw(screen *ebiten.Image) {
	if err := g.controller.Draw(screen); err != nil {
		log.Printf("controller draw: %v", err)
		return
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.context.LogicalWidth, g.context.LogicalHeight
}
