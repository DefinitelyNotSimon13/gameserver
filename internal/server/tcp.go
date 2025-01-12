package server

import (
	"github.com/DefinitelyNotSimon13/gameserver/internal/packet"
	"github.com/DefinitelyNotSimon13/gameserver/internal/processing"
	"log"
	"net"
)

// handleTCPConnection sends the client an Id and logs some messages.
func (s *Server) handleTCPConnection(conn net.Conn, id uint32) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("Connection %d closed or error: %v\n", id, err)
			break
		}

		log.Printf("Received %d bytes from client %d: %v\n", n, id, buf[:n])

		version, err := packet.ParsePacketVersion(buf[0])
		if err != nil {
			log.Printf("Failed to parse packet version: %v\n", err)
			continue
		}

		var responsePacket *packet.PacketV1
		switch version {
		case packet.VERSION_0:
			log.Println("Received deprecated version 0 packet")
		case packet.VERSION_1:
			responsePacket, err = processing.V1Packet(buf[:n])
		default:
			log.Println("Can't process packet, unknown version")
		}

		if err != nil {
			log.Printf("Failed to process V%d packet: %v\n", version, err)
		}

		if responsePacket == nil {
			log.Printf("Parsing of packet failed fatally, packet == nil.\n")
			log.Printf("Generating error packet for response...\n")
			responsePacket = packet.ErrorPacketV1()
		}

		response, err := responsePacket.ToBytes()
		if err != nil {
			log.Fatalf("Failed to convert packet into bytes: %v", err)
		}

		if _, err := conn.Write(response); err != nil {
			log.Printf("Failed to send response: %v\n", err)
		}
	}

	// bs := make([]byte, 4)
	// binary.LittleEndian.PutUint32(bs, id)
	//
	// log.Printf("Assigned TCP Id %d -> Writing bytes: %v\n", id, bs)
	// if _, err := conn.Write(bs); err != nil {
	// 	log.Printf("Error sending TCP message: %v\n", err)
	// 	return
	// }
	//
	// log.Printf("Sent greeting to %v\n", conn.RemoteAddr())
	// Here you could read from conn, handle protocol negotiation, etc.
}
