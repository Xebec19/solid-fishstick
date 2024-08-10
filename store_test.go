package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpicture"
	pathname := CASPathTransformFunc(key)
	fmt.Println(pathname)

	expectedPathname := "71056/ad8aa/24742/ea41e/a36fa/2e345/2a316/36e82"

	if pathname != expectedPathname {
		t.Errorf("have %s want %s", pathname, expectedPathname)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: DefaultPathTransformFunc,
	}
	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeStream("test", data); err != nil {
		t.Error(err)
	}
}
