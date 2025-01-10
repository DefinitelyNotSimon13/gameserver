package cli

import (
	"flag"
	"fmt"
)

type CliArgs struct {
	Address string
}

func ParseCli() *CliArgs {
	portPtr := flag.Int("port", 9000, "Port for the game")
	flag.Parse()

	return &CliArgs{
		Address: fmt.Sprintf("127.0.0.1:%d", *portPtr),
	}
}
