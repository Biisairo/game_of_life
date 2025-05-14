package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	gamemappb "github.com/Biisairo/game_of_life/proto/gamemap"
	lexiconpb "github.com/Biisairo/game_of_life/proto/lexicon"
	"google.golang.org/grpc"
)

var (
	port = flag.String("p", "08061", "Port")
)

func main() {
	flag.Parse()

	address := fmt.Sprintf("localhost:%v", port)

	// open tcp server
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Cannot open server: %v", err)
		return
	}
	defer lis.Close()

	// set grpc server
	server := grpc.NewServer(grpc.EmptyServerOption{})

	gameMapServer := struct {
		gamemappb.UnimplementedGamemapServer
	}{}

	gamemappb.RegisterGamemapServer(server, gameMapServer)

	lexiconServer := struct {
		lexiconpb.UnimplementedLexiconServer
	}{}

	lexiconpb.RegisterLexiconServer(server, lexiconServer)

	// run server
	server.Serve(lis)
}
