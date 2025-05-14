package main

import (
	"flag"
	"fmt"
	"log"

	gamemappb "github.com/Biisairo/game_of_life/proto/gamemap"
	lexiconpb "github.com/Biisairo/game_of_life/proto/lexicon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port = flag.String("p", "08061", "Port")
)

func main() {
	flag.Parse()

	address := fmt.Sprintf("localhost:%v", port)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		log.Fatalf("Fail to Create Client: %v", err)
	}
	defer conn.Close()

	gameMapClient := gamemappb.NewGamemapClient(conn)
	lexiconClient := lexiconpb.NewLexiconClient(conn)
}
