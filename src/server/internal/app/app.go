package app

import (
	"github.com/Biisairo/game_of_life/src/server/internal/engine"
	"github.com/Biisairo/game_of_life/src/server/server"
)

type App struct {
	gameMapServer *server.GameMapServer
	lexiconServer *server.LexiconServer

	engine *engine.Engine
}
