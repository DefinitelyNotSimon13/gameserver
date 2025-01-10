package server

import (
	"encoding/binary"
	"log"
	"net"
)

// handleTCPConnection sends the client an Id and logs some messages.
func (s *Server) handleTCPConnection(conn net.Conn, id uint32) {
	defer conn.Close()

	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, id)

	log.Printf("Assigned TCP Id %d -> Writing bytes: %v\n", id, bs)
	if _, err := conn.Write(bs); err != nil {
		log.Printf("Error sending TCP message: %v\n", err)
		return
	}

	log.Printf("Sent greeting to %v\n", conn.RemoteAddr())
	// Here you could read from conn, handle protocol negotiation, etc.
}
