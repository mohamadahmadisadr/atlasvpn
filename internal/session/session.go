package session

import "net"

type Session struct {
	ID         uint32
	RemoteAddr *net.UDPAddr
}
