package server

import (
	"fmt"
	"log"
	"net"
	"sync"
	"github.com/DefinitelyNotSimon13/raylib-c/server/internal/client"
)

type Server struct {
	addr        string
	tcpListener net.Listener
	udpConn     *net.UDPConn

	clients	map[uint32]*client.Client
	mu         sync.Mutex



	connectionId uint32
}

// NewServer creates a Server with default maps, locks, etc.
func NewServer(addr string) *Server {
	return &Server{
		addr:       addr,
		clients: make(map[uint32]*client.Client),
	}
}

// Start starts both TCP and UDP listeners and blocks until the server is shut down.
func (s *Server) Start() error {
	// Start TCP listener
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("error starting TCP server: %w", err)
	}
	s.tcpListener = listener
	log.Printf("TCP server is listening on %s\n", s.addr)

	// Start UDP listener
	udpAddr, err := net.ResolveUDPAddr("udp", s.addr)
	if err != nil {
		return fmt.Errorf("error resolving UDP addr: %w", err)
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return fmt.Errorf("error starting UDP server: %w", err)
	}
	s.udpConn = udpConn
	log.Printf("UDP server is listening on %s\n", s.addr)

	// Launch the TCP accept loop in a separate goroutine
	go s.acceptTCPConnections()

	// Start handling UDP in the current goroutine (blocking)
	s.handleUDPConnections()

	return nil
}

// acceptTCPConnections accepts incoming TCP connections in a loop.
func (s *Server) acceptTCPConnections() {
	defer s.tcpListener.Close()

	for {
		tcpConn, err := s.tcpListener.Accept()
		if err != nil {
			log.Printf("Error accepting TCP connection: %v\n", err)
			continue
		}

		// We increment the connectionId here for each new connection

		s.connectionId++
		s.clients[s.connectionId] = client.NewClient(s.connectionId, &tcpConn)

		// Handle each connection in its own goroutine
		go s.handleTCPConnection(tcpConn, s.connectionId)
		log.Printf("Currently %d clients connected.\n", len(s.clients))
	}
}

// handleUDPConnections is the main loop for reading incoming UDP packets.
func (s *Server) handleUDPConnections() {
	defer s.udpConn.Close()

	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := s.udpConn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("Error reading UDP data: %v\n", err)
			continue
		}
		if n < 1 {
			log.Println("Received fewer than 1 byte, can't read version")
			continue
		}

		version := buf[0] & 0x0F // Low nibble
		switch version {
		case PacketVersion0:
			s.handlePacketV0(buf, n, remoteAddr)
		default:
			log.Printf("Unknown version (%d) from %v\n", version, remoteAddr)
		}
	}
}
