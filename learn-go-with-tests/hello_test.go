package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloName(t *testing.T) {
	got := HelloName("Ariel")
	want := "Hello, Ariel"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
