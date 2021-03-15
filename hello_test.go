package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Thanh")
	want := "Hello, Thanh"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
