package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Pedro")
	want := "Hello, Pedro"

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
