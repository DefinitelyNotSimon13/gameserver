package processing

import (
	"github.com/DefinitelyNotSimon13/gameserver/internal/packet"
	"github.com/google/uuid"
	"log"
)

func processPlayerInit(p *packet.PacketV0) (*packet.PacketV0, error) {
	flags := packet.DefaultFlags()

	userName := string(p.Payload)
	log.Printf("Received UserName: %s\n", userName)

	return &packet.PacketV0{
		Version:    0,
		Type:       0,
		ClientId:   uuid.New(),
		Flags:      *flags,
		PayloadLen: 0,
	}, nil
}
