package udp

import (
	"fmt"
	"net"
	"time"
)

func RunClient() error {

	addr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 3000,
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return fmt.Errorf("Error dialing UDP: %v", err)
	}

	defer conn.Close()

	message := []byte("Hello from the udp client!")

	_, err = conn.Write(message)
	if err != nil {
		return fmt.Errorf("Error writing data: %v", err)
	}
	fmt.Println("📤 Message sent successfully!")

	err = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		fmt.Printf("Error setting deadline: %v", err)
	}

	buffer := make([]byte, MaxPacketSize)

	n, err := conn.Read(buffer)

	if err != nil {
		return fmt.Errorf("Error reading response: %v", err)
	}

	fmt.Printf("📥 Server response: %s\n", string(buffer[:n]))

	return nil

}
