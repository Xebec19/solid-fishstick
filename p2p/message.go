package p2p

import "net"

// RPC holds any arbitrary data over
// each transport between two nodes
type RPC struct {
	From    net.Addr
	Payload []byte
}
