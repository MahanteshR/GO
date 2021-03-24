package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Mahan")
		want := "Hello, Mahantesh"
		assertCorrectMessage(t, got, want)
	})

	t.Run("test by passing an empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, Mahan"
		assertCorrectMessage(t, got, want)
	})
}
