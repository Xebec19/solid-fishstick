package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	return NewStore(opts)
}

func tearDown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}

func TestPathTransformFunc(t *testing.T) {
	key := "bestpicture"
	pathkey := CASPathTransformFunc(key)
	fmt.Println(pathkey)

	expectedFilename := "71056ad8aa24742ea41ea36fa2e3452a31636e82"
	expectedPathname := "71056/ad8aa/24742/ea41e/a36fa/2e345/2a316/36e82"

	if pathkey.Pathname != expectedPathname {
		t.Errorf("have %s want %s", pathkey, expectedPathname)
	}

	if pathkey.Filename != expectedFilename {
		t.Errorf("have %s want %s", pathkey, expectedFilename)
	}
}

func TestDelete(t *testing.T) {

	s := newStore()
	key := "specials"
	data := []byte("some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {

	s := newStore()

	defer tearDown(t, s)

	for i := 0; i < 50; i++ {
		key := "123specials"
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

		if err := s.Delete(key); err != nil {
			t.Error(err)
		}
	}

}
