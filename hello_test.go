package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Mary", "")
		want := "Hello, Mary"
		assertCorrectMessage(t, got, want)
	})

	t.Run("print 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("print in Spanish", func(t *testing.T) {
		got := Hello("Sammy", "Spanish")
		want := "Hola, Sammy"
		assertCorrectMessage(t, got, want)
	})

	t.Run("print in French", func(t *testing.T) {
		got := Hello("Claire", "French")
		want := "Bonjour, Claire"
		assertCorrectMessage(t, got, want)
	})

	t.Run("print in Urdu", func(t *testing.T) {
		got := Hello("Baba", "Urdu")
		want := "Salam, Baba"
		assertCorrectMessage(t, got, want)
	})

}
