package server

import (
	"github.com/DefinitelyNotSimon13/gameserver/internal/client"
	"github.com/DefinitelyNotSimon13/gameserver/internal/packet"
	"log"
	"net"
)

func (s *Server) processPlayerInit(p *packet.PacketV1, tcpConn *net.Conn) (*packet.PacketV1, error) {
	flags := packet.DefaultFlags()

	userName := string(p.Payload)
	log.Printf("Received UserName: %s\n", userName)

	c := client.NewClient(userName, tcpConn)
	s.mu.Lock()
	s.clients[c.ClientId] = c
	log.Printf("1 Clients: %v\n", s.clients)
	s.mu.Unlock()

	return &packet.PacketV1{
		Version:    1,
		Type:       0,
		ClientId:   c.ClientId,
		Flags:      *flags,
		PayloadLen: 0,
	}, nil
}
