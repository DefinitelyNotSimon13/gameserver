package main

import (
	"github.com/DefinitelyNotSimon13/gameserver/internal/cli"
	"github.com/DefinitelyNotSimon13/gameserver/internal/server"
	"log"
)

func main() {

	cliArgs := cli.ParseCli()

	srv := server.NewServer(cliArgs.Address)

	if err := srv.Start(); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
