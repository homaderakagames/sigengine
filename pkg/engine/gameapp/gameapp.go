package gameapp

import (
	"fmt"
	"github.com/caarlos0/env/v10"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/homaderakagames/sigengine/pkg/engine/config"
	"github.com/homaderakagames/sigengine/pkg/engine/gameapp/game"
	"github.com/homaderakagames/sigengine/pkg/engine/gameapp/gamecontext"
	"github.com/homaderakagames/sigengine/pkg/engine/gameapp/input"
	"github.com/homaderakagames/sigengine/pkg/engine/gameapp/scene"
	"github.com/homaderakagames/sigengine/pkg/engine/resources"
	"log"
)

type GameApp struct {
	config          *config.Common
	input           *input.Input
	game            *game.Game
	context         *gamecontext.Context
	resourceManager *resources.ResourceManager
}

func NewGameApp(
	startingScene scene.Scene,
	resourceManager *resources.ResourceManager,
) *GameApp {
	ctx := gamecontext.NewContext()
	ctx.Put(resources.GetResourceManager, resourceManager)
	startingScene.OnLoad(ctx)

	ga := &GameApp{
		context:         ctx,
		game:            game.NewGame(startingScene, ctx),
		resourceManager: resourceManager,
	}
	ga.initInput()
	// TODO: parse and set window size ebiten.SetWindowSize(, 720)

	return ga
}

func (ga *GameApp) Run() {
	defer func() {
		// TODO: add save on exit
	}()

	if err := ebiten.RunGame(ga.game); err != nil {
		log.Printf("run game error: %v", err)
	}
}

func (ga *GameApp) ParseConfig(cfg interface{}) {
	if err := env.Parse(cfg); err != nil {
		panic(fmt.Sprintf("cannot parse config: %s", err))
	}
}

func (ga *GameApp) initInput() {
	ga.input = input.NewInput(&input.EbitenKeyboardAdapter{})

}

func (ga *GameApp) initCommon() {
	var cfg config.Common
	ga.ParseConfig(&cfg)
	ga.config = &cfg
}
