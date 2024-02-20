package depinj

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := &bytes.Buffer{}
	Greet(buffer, "Elodie")
	got := buffer.String()
	want := "Hello, Elodie\n"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
