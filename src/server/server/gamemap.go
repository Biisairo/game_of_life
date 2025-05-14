package server

import (
	"context"

	commonpb "github.com/Biisairo/game_of_life/proto/common"
	gamemappb "github.com/Biisairo/game_of_life/proto/gamemap"
	"github.com/Biisairo/game_of_life/src/server/internal/engine"
)

// rpc SyncCheck(SyncCheckRequest) returns (SyncCheckResponse);
// rpc SyncGameMap(common.Empty) returns (SyncGameMapResponse);
// rpc NewPattern(NewPatternRequest) returns (common.Empty);
// rpc StreamMap(common.Empty) returns (stream StreamMapResponse);

type GameMapServer struct {
	gamemappb.UnimplementedGamemapServer

	gameMap *engine.Game
}

func NewGameMapServer() *GameMapServer {
	s := &GameMapServer{}
	return s
}

func (s *GameMapServer) SyncCheck(ctx context.Context, record *gamemappb.SyncCheckRequest) (*gamemappb.SyncCheckResponse, error) {
}

func (s *GameMapServer) SyncGameMap(ctx context.Context, record *commonpb.Empty) (*gamemappb.SyncGameMapResponse, error) {
}

func (s *GameMapServer) NewPattern(ctx context.Context, record *gamemappb.NewPatternRequest) (*commonpb.Empty, error) {
}

func (s *GameMapServer) StreamMap(_ *commonpb.Empty, stream gamemappb.Gamemap_StreamMapServer) error {
}
