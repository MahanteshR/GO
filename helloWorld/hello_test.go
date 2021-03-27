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
		got := Hello("Mahantesh", "")
		want := "Hello, Mahantesh"
		assertCorrectMessage(t, got, want)
	})

	t.Run("test by passing an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, Mahan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("spanish added", func(t *testing.T) {
		got := Hello("Elodie", "spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("French test", func(t *testing.T) {
		got := Hello("", "french")
		want := "Bonjour, Mahan"
		assertCorrectMessage(t, got, want)
	})
}
