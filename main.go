package main

import (
	"fmt"
	"log"

	"github.com/Xebec19/solid-fishstick/p2p"
)

func OnPeer(peer p2p.Peer) error {
	return fmt.Errorf("failed the onpeer func")
}

func main() {

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: func(a any) error { return nil },
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%v", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
