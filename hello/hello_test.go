package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello with an empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}