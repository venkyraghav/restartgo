package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Chris"), "Hello, Chris")
	})
	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		assertCorrectMessage(t, Hello(""), "Hello, World")
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
