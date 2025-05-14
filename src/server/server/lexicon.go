package server

import (
	"context"

	commonpb "github.com/Biisairo/game_of_life/proto/common"
	lexiconpb "github.com/Biisairo/game_of_life/proto/lexicon"
	"github.com/Biisairo/game_of_life/src/server/internal/engine"
)

type LexiconServer struct {
	lexiconpb.UnimplementedLexiconServer

	lexicon *engine.Lexicon
}

func NewLexiconServer() *LexiconServer {
	s := &LexiconServer{}
	return s
}

func (s *LexiconServer) SaveRecord(ctx context.Context, record *lexiconpb.Record) (*commonpb.Empty, error) {
}

func (s *LexiconServer) GetLexicon(_ *commonpb.Empty, stream lexiconpb.Lexicon_GetLexiconServer) error {
}
