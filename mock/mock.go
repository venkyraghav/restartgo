package mock

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (c *ConfigurableSleeper) SetDuration(duration time.Duration) {
	c.duration = duration
}

func (c *ConfigurableSleeper) SetSleep(sleep func(time.Duration)) {
	c.sleep = sleep
}

func Countdown(writer io.Writer, sleeper Sleeper, numberToCount int) {
	for i := numberToCount; i > 0; i-- {
		fmt.Fprintf(writer, "%d\n", i)
		sleeper.Sleep()
	}
	fmt.Fprintf(writer, "Go!")
}
