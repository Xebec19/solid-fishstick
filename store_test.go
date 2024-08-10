package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpicture"
	pathkey := CASPathTransformFunc(key)
	fmt.Println(pathkey)

	originalPathname := "71056ad8aa24742ea41ea36fa2e3452a31636e82"
	expectedPathname := "71056/ad8aa/24742/ea41e/a36fa/2e345/2a316/36e82"

	if pathkey.Pathname != expectedPathname {
		t.Errorf("have %s want %s", pathkey, expectedPathname)
	}

	if pathkey.Filename != originalPathname {
		t.Errorf("have %s want %s", pathkey, originalPathname)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)

	key := "specials"
	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := ioutil.ReadAll(r)

	if string(b) != string(data) {
		t.Errorf("want %s have %s", data, b)
	}
}
