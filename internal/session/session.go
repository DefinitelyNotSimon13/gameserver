package session

import (
	"github.com/DefinitelyNotSimon13/gameserver/internal/client"
	"github.com/google/uuid"
	"math/rand"
	"sync"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const length = 6

type Session struct {
	Token       string
	Connections map[uuid.UUID]*client.Client
	mu          sync.Mutex
}

func NewSession() *Session {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return &Session{
		Token:       string(b),
		Connections: make(map[uuid.UUID]*client.Client),
	}
}

func (s *Session) AddClient(c *client.Client) {
	c.ConnectedSession = &s.Token

	s.Connections[c.ClientId] = c
}

func (s *Session) RemoveClient(c *client.Client) {
	c.ConnectedSession = nil

	delete(s.Connections, c.ClientId)

}

func (s *Session) InvalidateSession() {
	for _, c := range s.Connections {
		c.ConnectedSession = nil
	}
	s.Connections = nil
}
