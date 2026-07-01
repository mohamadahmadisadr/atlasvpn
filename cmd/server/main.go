package main

import (
	"log"

	"github.com/mohamadahmadisadr/atlasvpn/internal/transport/udp"
)

func main() {

	if err := udp.Serve(); err != nil {
		log.Fatal(err)
	}
}
