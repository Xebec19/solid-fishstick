package p2p

import (
	"encoding/gob"
	"io"
)

type GOBDecoder struct{}

type Decoder interface {
	Decode(io.Reader, any) error
}

func (doc GOBDecoder) Decode(r io.Reader, v any) error {
	return gob.NewDecoder(r).Decode(v)
}

type GODDecoder struct{}

func (dec GODDecoder) Decode(r io.Reader, v any) error {
	return gob.NewDecoder(r).Decode(v)
}
