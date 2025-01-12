package client

import (
	"github.com/google/uuid"
	"net"
)

type Client struct {
	ClientId         uuid.UUID
	ConnectedSession *string
	Username         string
	UDPAddr          *net.UDPAddr
	TCPConn          *net.Conn
}

func NewClient(username string, tcpConn *net.Conn) *Client {
	return &Client{
		ClientId: uuid.New(),
		TCPConn:  tcpConn,
	}
}
