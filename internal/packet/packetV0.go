package packet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"math"
)

const (
	HeaderSize  = 23
	ClientIdLen = 16
)

type PacketV0 struct {
	Version    uint8     // Byte 0
	Type       uint8     // Byte 1
	ClientId   uuid.UUID // Byte 2 -17
	Flags      Flags     // Byte 18
	PayloadLen uint32
	Payload    []byte
	//TODO Checksum - needed? Not really via TCP right? Maybe for udp
}

func (p *PacketV0) ToBytes() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Grow(int(HeaderSize + p.PayloadLen))

	if err := buf.WriteByte(p.Version); err != nil {
		return nil, err
	}

	if err := buf.WriteByte(p.Type); err != nil {
		return nil, err
	}

	idBytes, err := p.ClientId.MarshalBinary()

	if err != nil {
		return nil, err
	}

	if _, err := buf.Write(idBytes); err != nil {
		return nil, err
	}

	if err := buf.WriteByte(p.Flags.ToByte()); err != nil {
		return nil, err
	}

	payloadLen := uint32(len(p.Payload))
	if payloadLen != p.PayloadLen {
		return nil, fmt.Errorf("payloadLen was not set correctly - got: %d, expected: %d", p.PayloadLen, payloadLen)
	}

	if err := binary.Write(buf, binary.LittleEndian, payloadLen); err != nil {
		return nil, err
	}

	if _, err := buf.Write(p.Payload); err != nil {
		return nil, err
	}

	if buf.Len() != int(HeaderSize+p.PayloadLen) {
		return nil, fmt.Errorf("created buffer has incorrect length- has: %d, exptected: %d", buf.Len(), int(HeaderSize+p.PayloadLen))
	}

	return buf.Bytes(), nil

}

func ErrorPacketV0() *PacketV0 {
	return &PacketV0{
		Version:    0,
		Type:       255,
		Flags:      *DefaultFlags(),
		PayloadLen: 0,
		Payload:    make([]byte, 0),
	}
}

func ParsePacketV0(data []byte) (*PacketV0, error) {
	log.Println("Parsing Packet...")
	if len(data) < HeaderSize {
		return nil, errors.New("packet to short")
	}

	packetType, err := parsePacketType(data[1])
	if err != nil {
		return nil, errors.New("invalid packet type")
	}

	clientId, err := uuid.FromBytes(data[2 : 2+ClientIdLen])
	if err != nil {
		return nil, errors.Join(errors.New("failed to parse client id"), err)
	}

	flags := parseFlags(data[18])

	payloadLen := binary.LittleEndian.Uint32(data[19:23])
	if len(data) < HeaderSize+int(payloadLen) {
		return nil, fmt.Errorf("packet payload length mismatch: declared %d, but data is %d bytes", payloadLen, len(data)-HeaderSize)
	}

	payload := data[23:(23 + payloadLen)]

	log.Printf("\tPayload: %v", payload)
	log.Printf("\tClient Id: %v\n", clientId)
	log.Printf("\tFlags: %v\n", flags)
	log.Printf("\tPayloadLength: %d", payloadLen)

	return &PacketV0{
		Version:    0,
		Type:       packetType,
		ClientId:   clientId,
		Flags:      flags,
		PayloadLen: payloadLen,
		Payload:    payload,
	}, nil
}

// parseSenderId extracts the sender Id from bytes 1..4 (little-endian).
func ParseSenderId(data []byte) uint32 {
	fmt.Printf("Fun")
	return binary.LittleEndian.Uint32(data[1:5])
}

// parseCoordinates extracts x, y, z from the data (bytes 5..8, 9..12, 13..16).
func ParseCoordinates(data []byte) (float32, float32, float32) {
	xBits := binary.LittleEndian.Uint32(data[5:9])
	yBits := binary.LittleEndian.Uint32(data[9:13])
	zBits := binary.LittleEndian.Uint32(data[13:17])
	return math.Float32frombits(xBits),
		math.Float32frombits(yBits),
		math.Float32frombits(zBits)
}
