package mock

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type SpyCountdownOperations struct {
	Calls   []string
	Buffer  *bytes.Buffer
	Sleeper Sleeper
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
	s.Sleeper.Sleep()
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)

	return s.Buffer.Write(p)
}

type SpyTime struct {
	durationSlept []time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = append(s.durationSlept, duration)
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to 1 with newline after every number", func(t *testing.T) {
		spyTime := &SpyTime{}
		sleeper := &ConfigurableSleeper{1 * time.Second, spyTime.Sleep}
		buffer := &SpyCountdownOperations{Buffer: &bytes.Buffer{}, Sleeper: sleeper}
		Countdown(buffer, buffer, 3)
		got := buffer.Buffer.String()
		want := "3\n2\n1\nGo!"
		got2 := buffer.Calls
		want2 := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		if !reflect.DeepEqual(got2, want2) {
			t.Errorf("got %q want %q", got2, want2)
		}
		got3 := spyTime.durationSlept
		want3 := []time.Duration{
			1 * time.Second,
			1 * time.Second,
			1 * time.Second,
		}
		if !reflect.DeepEqual(got3, want3) {
			t.Errorf("slept for %q want %q", got3, want3)
		}
	})
}
