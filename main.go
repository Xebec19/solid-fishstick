package main

import (
	"log"

	"github.com/Xebec19/solid-fishstick/p2p"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {

	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}

	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{
		// ListenAddr:        ":3000",
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}

	return NewFileServer(fileServerOpts)
}

func main() {

	s1 := makeServer(":3000", "")

	s2 := makeServer(":4000", ":3000")

	go func() { log.Fatal(s1.Start()) }()

	go func() { log.Fatal(s2.Start()) }()

	// go func() {
	// 	time.Sleep(time.Second * 3)
	// 	s1.Stop()
	// 	s2.Stop()
	// }()
}
