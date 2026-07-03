package main

import (
	"fmt"
	"log"

	"github.com/mohamadahmadisadr/atlasvpn/internal/transport/udp"
)

func main() {

	if err := udp.Serve(); err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}
}
