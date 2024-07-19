package p2p

type Handshaker interface {
	Handshake() error
}

type HandshakerFunc func(any) error
