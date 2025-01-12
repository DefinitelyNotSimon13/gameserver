package packet

import (
	"errors"
)

const (
	VERSION_0 = 0
	VERSION_1 = 1
)

func ParsePacketVersion(version byte) (uint8, error) {
	switch version {
	case VERSION_0:
		return 0, nil
	case VERSION_1:
		return 1, nil
	default:
		return 0, errors.New("invalid version")
	}
}
