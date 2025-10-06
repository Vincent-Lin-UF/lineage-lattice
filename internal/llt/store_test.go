package llt

import "testing"

func TestStore_PutGet(t *testing.T) {
	s := NewStore()
	s.Put("Hello", []byte("World"))

	got, ok := s.Get("Hello")
	if !ok {
		t.Fatalf("Expected key to exist")
	}

	if string(got) != "World" {
		t.Fatalf("Got %q, Want %q", string(got), "World")
	}
}
