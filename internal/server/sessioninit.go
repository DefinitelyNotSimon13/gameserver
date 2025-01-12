package server

import (
	"errors"
	"fmt"
	"github.com/DefinitelyNotSimon13/gameserver/internal/packet"
	"github.com/DefinitelyNotSimon13/gameserver/internal/session"
	"log"
)

func (s *Server) processSessionInit(p *packet.PacketV1) (*packet.PacketV1, error) {
	flags := packet.DefaultFlags()

	var sessionToken string
	if p.PayloadLen == 0 {

		sess := session.NewSession()
		s.mu.Lock()
		s.sessions[sess.Token] = sess
		s.mu.Unlock()
		sessionToken = sess.Token
		log.Printf("Created new session with token %s\n", sessionToken)
	} else {
		sessionToken = string(p.Payload)
		log.Printf("Received session token: %s\n", sessionToken)
		if _, ok := s.sessions[sessionToken]; !ok {
			log.Printf("Session with token %s does not exist\n", sessionToken)
			return nil, errors.New("session does not exist")
		}
	}

	s.mu.Lock()
	s.sessions[sessionToken].AddClient(s.clients[p.ClientId])
	for token, sess := range s.sessions {
		fmt.Printf("Session Token: %s\n", token)
		for clientId, client := range sess.Connections {
			fmt.Printf("\tClient ID: %s, Username: %s\n", clientId, client.Username)
		}
	}
	s.mu.Unlock()

	return &packet.PacketV1{
		Version:    1,
		Type:       0,
		ClientId:   p.ClientId,
		Flags:      *flags,
		PayloadLen: 0,
	}, nil
}
