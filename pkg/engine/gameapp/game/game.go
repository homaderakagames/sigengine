package game

import (
	"context"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/homaderakagames/sigengine/pkg/engine/gameapp/gamecontext"
	"github.com/homaderakagames/sigengine/pkg/engine/gameapp/scene"
	"log"
)

type Game struct {
	currentScene   scene.Scene
	curSceneParams *scene.Params

	context *gamecontext.Context
	cancel  context.CancelFunc
}

func NewGame(
	scene scene.Scene,
	ctx *gamecontext.Context,
	cancel context.CancelFunc,
) *Game {

	return &Game{
		currentScene:   scene,
		context:        ctx,
		cancel:         cancel,
		curSceneParams: &scene.Params{},
	}
}

func (g *Game) Update() error {
	select {
	case <-g.context.Done():
		return ebiten.Termination
	default:
	}

	if ok := g.curSceneParams.Get(g.context); ok {
		g.currentScene = g.curSceneParams.Scene
	}

	return g.currentScene.Update(g.context)
}

func (g *Game) Draw(screen *ebiten.Image) {
	if err := g.currentScene.Draw(screen); err != nil {
		log.Printf("current scene draw: %v", err)
		return
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.context.LogicalWidth, g.context.LogicalHeight
}
