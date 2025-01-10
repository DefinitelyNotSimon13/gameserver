package main

import (
	"github.com/DefinitelyNotSimon13/gameserver/internal/server"
	"log"
)

func main() {
	addr := "localhost:9000"

	srv := server.NewServer(addr)

	if err := srv.Start(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
