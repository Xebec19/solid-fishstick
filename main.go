package main

import (
	"log"

	"github.com/Xebec19/solid-fishstick/p2p"
)

func main() {

	tr := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: func() { return },
	}

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
