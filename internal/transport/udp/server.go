package udp

import (
	"fmt"
	"net"

	"github.com/mohamadahmadisadr/atlasvpn/internal/protocol"
	"github.com/mohamadahmadisadr/atlasvpn/internal/session"
)

func Serve() error {
	sessionManager := session.NewManager()
	addr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 3000,
	}

	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		return fmt.Errorf("Failed to bind to UDP: %v", err)
	}

	defer conn.Close()

	fmt.Printf("UDP server listening on %s\n", conn.LocalAddr().String())

	buffer := make([]byte, MaxPacketSize)

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading packet: %v", err)
			continue
		}

		pckt, err := protocol.DeserializePacket(buffer[:])
		if err != nil {
			return err
		}

		switch pckt.Type {
		case protocol.PacketTypeHandshake:
			sessionID := session.GenerateSessionID()
			session := sessionManager.Create(sessionID, remoteAddr)
			response := protocol.Packet{
				Version:  1,
				Type:     protocol.PacketTypeHandshake,
				PacketID: session.ID,
				Payload:  []byte("OK"),
			}

			data, err := response.SerializePacket()
			if err != nil {
				return err
			}
			if _, err := conn.WriteToUDP(data, session.RemoteAddr); err != nil {
				return err
			}
		}
		message := string(buffer[:n])
		fmt.Printf("Received %d bytes from %s: %s\n", n, remoteAddr, message)

		response := []byte("Message received!")

		_, err = conn.WriteToUDP(response, remoteAddr)
		if err != nil {
			return fmt.Errorf("Error sending response to %s: %v", remoteAddr, err)
		}

	}

}
