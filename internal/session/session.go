package session

import (
	"github.com/DefinitelyNotSimon13/gameserver/internal/client"
	"github.com/google/uuid"
	"sync"
)

type Session struct {
	SessionId   uuid.UUID
	Connections map[uuid.UUID]*client.Client
	mu          sync.Mutex
}

func NewSession() (*Session, error) {
	id := uuid.New()

	return &Session{
		SessionId:   id,
		Connections: make(map[uuid.UUID]*client.Client, 0),
	}, nil
}

func (s *Session) AddClient(c *client.Client) {
	c.ConnectedSession = &s.SessionId

	s.Connections[c.ClientId] = c

}

func (s *Session) RemoveClient(c *client.Client) {
	c.ConnectedSession = nil

	delete(s.Connections, c.ClientId)

}

func (s *Session) InvalidateSession() {
	for _, client := range s.Connections {
		client.ConnectedSession = nil
	}
	s.Connections = nil
}
