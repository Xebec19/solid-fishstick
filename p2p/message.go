package p2p

// Message holds any arbitrary data over
// each transport between two nodes
type Message struct {
	Payload []byte
}
