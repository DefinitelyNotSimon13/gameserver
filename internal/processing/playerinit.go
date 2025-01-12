package processing

import (
	"github.com/DefinitelyNotSimon13/gameserver/internal/packet"
	"github.com/google/uuid"
	"log"
)

func processPlayerInit(p *packet.PacketV1) (*packet.PacketV1, error) {
	flags := packet.DefaultFlags()

	userName := string(p.Payload)
	log.Printf("Received UserName: %s\n", userName)

	return &packet.PacketV1{
		Version:    1,
		Type:       0,
		ClientId:   uuid.New(),
		Flags:      *flags,
		PayloadLen: 0,
	}, nil
}
