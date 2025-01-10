package client

import (
	"github.com/google/uuid"
	"net"
)

type Client struct {
	ClientId         uuid.UUID
	ConnectedSession *uuid.UUID
	Username         string
	UDPAddr          *net.UDPAddr
	TCPConn          *net.Conn
}

func NewClient(id uuid.UUID, tcpConn *net.Conn) *Client {
	return &Client{
		ClientId: id,
		TCPConn:  tcpConn,
	}
}
