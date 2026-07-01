package udp

import (
	"fmt"
	"net"
)

func Serve() error {
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

		message := string(buffer[:n])
		fmt.Printf("Received %d bytes from %s: %s\n", n, remoteAddr, message)

		response := []byte("Message received!")

		_, err = conn.WriteToUDP(response, remoteAddr)
		if err != nil {
			return fmt.Errorf("Error sending response to %s: %v", remoteAddr, err)
		}

	}

}
