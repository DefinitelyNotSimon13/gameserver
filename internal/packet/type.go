package packet

import (
	"errors"
)

const (
	PLAYER_INIT     = 0
	SESSION_INIT    = 1
	TYPE_3          = 2
	TYPE_4          = 3
	TYPE_5          = 4
	TYPE_6          = 5
	TYPE_7          = 6
	PLAYER_POSITION = 7
)

func parsePacketType(data byte) (uint8, error) {
	switch data {
	case PLAYER_INIT:
		return PLAYER_INIT, nil
	case SESSION_INIT:
		return SESSION_INIT, nil
	case PLAYER_POSITION:
		return PLAYER_POSITION, nil
	default:
		return 0, errors.New("invalid type")
	}
}
