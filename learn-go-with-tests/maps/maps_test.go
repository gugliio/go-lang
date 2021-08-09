package main

import "testing"

func TestSearches(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}
	got, _ := dictionary.Search("test")
	want := "this is a test"
	assertString(t, got, want)

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrorNotFound.Error()

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertString(t, err.Error(), want)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("se obtuvo %s se esperaba %s, dado %s", got, want, "test")
	}
}

func assertError(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("se obtuvo %s se esperaba %s, dado %s", got, want, "test")
	}
}
