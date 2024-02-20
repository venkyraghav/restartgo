package mock

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to 1 with newline after every number", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer)
		got := buffer.String()
		want := "3\n2\n1\n"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
