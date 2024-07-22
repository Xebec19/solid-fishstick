package p2p

import "net"

// Message holds any arbitrary data over
// each transport between two nodes
type Message struct {
	From    net.Addr
	Payload []byte
}
