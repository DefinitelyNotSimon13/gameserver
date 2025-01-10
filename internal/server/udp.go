package server

import (
	"log"
	"net"
)

// handlePacketV0 is a method for handling version-0 UDP packets.
func (s *Server) handlePacketV0(data []byte, length int, remoteAddr *net.UDPAddr) {
	// Safe to assume we've already checked length >= 1
	if length < 17 {
		log.Println("Received too few bytes for PacketVersion0 (need >= 17)")
		return
	}

	senderId := parseSenderId(data)
	x, y, z := parseCoordinates(data)

	log.Printf(
		"V0 packet from %v with Id %d => x=%.2f, y=%.2f, z=%.2f\n",
		remoteAddr, senderId, x, y, z,
	)

	// Register the new client if it doesn't exist
	s.mu.Lock()
	log.Printf("Trying to access client with id %d\n", senderId)

	// Gotta check if client exists at all
	if s.clients[senderId].UDPAddr == nil {
		s.clients[senderId].UDPAddr = remoteAddr
		log.Printf("Registered new UDP client with Id %d at %v\n", senderId, remoteAddr)
	}
	// Broadcast to all other clients
	for clientId, client := range s.clients {
		if clientId == senderId {
			continue
		}
		if _, err := s.udpConn.WriteToUDP(data[:length], client.UDPAddr); err != nil {
			log.Printf("Error broadcasting to client %d at %v: %v\n", client.Id, client.UDPAddr, err)
		}
	}
	s.mu.Unlock()
}
