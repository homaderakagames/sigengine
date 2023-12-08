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
	"log"
)

type GameApp struct {
	config  *config.Common
	input   *input.Input
	game    *game.Game
	context *gamecontext.Context
}

func NewGameApp(ctrl scene.Scene) *GameApp {
	ctx := gamecontext.NewContext()
	ctrl.OnLoad(ctx)

	ga := &GameApp{
		context: ctx,
		game:    game.NewGame(ctrl, ctx),
	}
	ga.initInput()

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
	ga.input = input.NewInput()

}

func (ga *GameApp) initCommon() {
	var cfg config.Common
	ga.ParseConfig(&cfg)
	ga.config = &cfg
}
