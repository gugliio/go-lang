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

	// For helper functions, it's a good idea to accept a testing.TB which is an interface that *testing.T
	// and *testing.B both satisfy, so you can call helper functions from a test, or a benchmark.

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say Hello with name", func(t *testing.T) {
		got := HelloName("Ariel", "")
		want := "Hello, Ariel"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say Hello without name", func(t *testing.T) {
		got := HelloName("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := HelloName("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := HelloName("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})

}
