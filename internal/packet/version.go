package packet

import (
	"errors"
)

const (
	VERSION_0 = 0
)

func ParsePacketVersion(version byte) (uint8, error) {
	switch version {
	case VERSION_0:
		return 0, nil
	default:
		return 0, errors.New("invalid version")
	}
}
