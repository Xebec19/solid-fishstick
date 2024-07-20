package main

import (
	"log"

	"github.com/Xebec19/solid-fishstick/p2p"
)

func noopFunc(a int32) error {
	return nil
}

func main() {

	tr := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: noopFunc(12),
	}

	conn := p2p.NewTCPTransport(tr)

	if err := conn.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
