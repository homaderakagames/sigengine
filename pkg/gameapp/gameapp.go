package gameapp

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v10"
	"github.com/homaderakagames/sigengine/pkg/config"
	"github.com/homaderakagames/sigengine/pkg/game"
	"github.com/homaderakagames/sigengine/pkg/gameapp/input"
)

type GameApp struct {
	config *config.Common
	input  *input.Input

	context *game.Context
	cancel  context.CancelFunc
}

func NewGameApp() *GameApp {
	ctx, cancel := game.NewContext(context.Background()).WithCancel()
	ga := &GameApp{
		context: ctx,
		cancel:  cancel,
	}
	ga.ParseConfig()
	ga.initInput()

	return ga
}

func (ga *GameApp) ParseConfig() {
	var cfg config.Common
	if err := env.Parse(&cfg); err != nil {
		panic(fmt.Sprintf("parse config: %v", err))
	}
	ga.config = &cfg
}

func (ga *GameApp) initInput() {
	ga.input = input.NewInput()
}
