package main

import (
	"log"

	"github.com/Xebec19/solid-fishstick/p2p"
)

func main() {

	tr := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: func(a any) error { return nil },
		Decoder:       p2p.DefaultDecoder{},
	}

	conn := p2p.NewTCPTransport(tr)

	if err := conn.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
