package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people in English", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Chris", English), "Hello, Chris")
	})
	t.Run("say 'Hello, World!' in English when an empty string is supplied", func(t *testing.T) {
		assertCorrectMessage(t, Hello("", English), "Hello, World")
	})
	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Chris", Spanish), "Hola, Chris")
	})
	t.Run("say 'Hello, World!' in Spanish when an empty string is supplied", func(t *testing.T) {
		assertCorrectMessage(t, Hello("", Spanish), "Hola, World")
	})
	t.Run("saying hello to people in French", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Chris", French), "Bonjour, Chris")
	})
	t.Run("say 'Hello, World!' in French when an empty string is supplied", func(t *testing.T) {
		assertCorrectMessage(t, Hello("", French), "Bonjour, World")
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
