package client

import (
	"net"
)

type Client struct {
	Id      uint32
	UDPAddr *net.UDPAddr
	// Pointer or not?
	TCPConn *net.Conn
	// Status?
}

func NewClient(id uint32, tcpConn *net.Conn) *Client {
	return &Client{
		Id:      id,
		TCPConn: tcpConn,
	}
}
