package session

import (
	"github.com/DefinitelyNotSimon13/gameserver/internal/client"
	"github.com/google/uuid"
)

type Session struct {
	SessionId   uuid.UUID
	ConnClients []client.Client
}
