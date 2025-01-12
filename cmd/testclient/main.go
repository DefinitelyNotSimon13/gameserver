package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net"
	"os"
	"strings"
)

const (
	SGCPVersion = 1
)

func main() {
	// CLI flags
	protocol := flag.String("protocol", "tcp", "Protocol to use (tcp or udp)")
	serverAddr := flag.String("address", "127.0.0.1:9000", "Server address")
	flag.Parse()

	switch *protocol {
	case "tcp":
		handleTCP(*serverAddr)
	case "udp":
		handleUDP(*serverAddr)
	default:
		log.Fatalf("Unsupported protocol: %s", *protocol)
	}
}

func handleTCP(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Error connecting to TCP server: %v", err)
	}
	defer conn.Close()

	log.Println("Connected to TCP server. Type 'exit' to close the connection.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter message type (or 'exit' to quit): ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			break
		}

		// Parse message type
		var msgType uint8
		_, err := fmt.Sscanf(input, "%d", &msgType)
		if err != nil {
			log.Println("Invalid message type. Please enter a number.")
			continue
		}

		fmt.Print("Enter payload (or leave empty): ")
		scanner.Scan()
		payload := scanner.Text()

		// Generate a random UserId for testing
		userId := uuid.New()

		userIdBytes, err := userId.MarshalBinary()
		userIdBytes = make([]byte, 16)
		if err != nil {
			log.Printf("Error retrieving UUID bytes: %v\n", err)
			continue
		}

		// Craft and send the message
		msg, err := craftMessage(msgType, userIdBytes, []byte(payload))
		if err != nil {
			log.Printf("Error crafting message: %v", err)
			continue
		}

		if _, err := conn.Write(msg); err != nil {
			log.Printf("Error sending TCP message: %v\n", err)
			break
		}

		log.Println("Message sent. Waiting for response...")

		// Read response
		response := make([]byte, 1024)
		n, err := conn.Read(response)
		if err != nil {
			log.Printf("Error reading response: %v\n", err)
			break
		}

		log.Printf("Response: %v\n", response[:n])
	}
}

func handleUDP(address string) {
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatalf("Error resolving UDP address: %v", err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("Error connecting to UDP server: %v", err)
	}
	defer conn.Close()

	log.Println("Connected to UDP server. Type 'exit' to close the connection.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter message type (or 'exit' to quit): ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			break
		}

		// Parse message type
		var msgType uint8
		_, err := fmt.Sscanf(input, "%d", &msgType)
		if err != nil {
			log.Println("Invalid message type. Please enter a number.")
			continue
		}

		fmt.Print("Enter payload (or leave empty): ")
		scanner.Scan()
		payload := scanner.Text()

		// Generate a random UserId for testing
		userId := uuid.New()
		if err != nil {
			log.Printf("Error creating UUID: %v\n", err)
			continue
		}

		userIdBytes, err := userId.MarshalBinary()
		if err != nil {
			log.Printf("Error retrieving UUID bytes: %v\n", err)
			continue
		}

		// Craft and send the message
		msg, err := craftMessage(msgType, userIdBytes, []byte(payload))
		if err != nil {
			log.Printf("Error crafting message: %v", err)
			continue
		}

		if _, err := conn.Write(msg); err != nil {
			log.Printf("Error sending UDP message: %v\n", err)
			break
		}

		log.Println("Message sent.")
	}
}

func craftMessage(msgType uint8, userId, payload []byte) ([]byte, error) {
	buf := &bytes.Buffer{}

	// Write SGCP Version
	if err := buf.WriteByte(SGCPVersion); err != nil {
		return nil, err
	}

	// Write Message Type
	if err := buf.WriteByte(msgType); err != nil {
		return nil, err
	}

	// Write UserId (16 bytes)
	if len(userId) != 16 {
		return nil, fmt.Errorf("UserId must be 16 bytes")
	}
	if _, err := buf.Write(userId); err != nil {
		return nil, err
	}

	// Write Flags (1 byte)
	if err := buf.WriteByte(0); err != nil {
		return nil, err
	}

	// Write Payload Length (4 bytes, little-endian)
	payloadLength := uint32(len(payload))
	if err := binary.Write(buf, binary.LittleEndian, payloadLength); err != nil {
		return nil, err
	}

	// Write Payload
	if _, err := buf.Write(payload); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
