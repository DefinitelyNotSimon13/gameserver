package server

import (
	"errors"
	"github.com/DefinitelyNotSimon13/gameserver/internal/packet"
	"log"
	"net"
)

func (s *Server) processV1Packet(buf []byte, tcpConn *net.Conn) (*packet.PacketV1, error) {
	log.Printf("Processing V1 Packet with length %d\n", len(buf))
	p, err := packet.ParsePacketV1(buf)
	if err != nil {
		return nil, errors.Join(errors.New("Failed to process packet"), err)
	}
	log.Printf("Successfully parsed packet: %v\n", p)

	switch p.Type {
	case packet.PLAYER_INIT:
		log.Println("Type is: PlayerInit Packet")
		return s.processPlayerInit(p, tcpConn)

	case packet.SESSION_INIT:
		log.Println("Type is: SessionInit Packet")
		return s.processSessionInit(p)

	case packet.PLAYER_POSITION:
		log.Println("Type is: PlayerPosition Packet")

	default:
		return nil, errors.New("unknown packet type")
	}

	return nil, nil
}
