package server

import (
	"fmt"
	"github.com/DefinitelyNotSimon13/gameserver/internal/packet"
	"github.com/DefinitelyNotSimon13/gameserver/internal/session"
	"github.com/google/uuid"
	"log"
	"net"
	"sync"
)

type Server struct {
	addr        string
	tcpListener net.Listener
	udpConn     *net.UDPConn

	sessions map[uuid.UUID]*session.Session
	mu       sync.Mutex

	connectionId uint32
}

// NewServer creates a Server with default maps, locks, etc.
func NewServer(addr string) *Server {
	return &Server{
		addr:     addr,
		sessions: make(map[uuid.UUID]*session.Session),
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
		// s.clients[s.connectionId] = client.NewClient(uuid.New(), &tcpConn)

		// Handle each connection in its own goroutine
		go s.handleTCPConnection(tcpConn, s.connectionId)
		// log.Printf("Currently %d clients connected.\n", len(s.clients))
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

		version, err := packet.ParsePacketVersion(buf[0])
		if err != nil {
			log.Printf("Failed to parse packet version: %v\n", err)
			continue
		}
		switch version {
		case packet.VERSION_0:
			s.handlePacketV1(buf, n, remoteAddr)
		default:
			log.Printf("Unknown version (%d) from %v\n", version, remoteAddr)
		}
	}
}
