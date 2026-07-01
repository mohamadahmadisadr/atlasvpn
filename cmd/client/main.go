package main

import (
	"log"

	"github.com/mohamadahmadisadr/atlasvpn/internal/transport/udp"
)

func main() {
	if err := udp.RunClient(); err != nil {
		log.Fatalf("failed send request")
	}
}
