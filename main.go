package main

import (
	"os"
	"time"

	"github.com/venkyraghav/restartgo/mock"
)

func main() {
	sleeper := &mock.ConfigurableSleeper{}
	sleeper.SetDuration(1 * time.Second)
	sleeper.SetSleep(time.Sleep)
	mock.Countdown(os.Stdout, sleeper, 5)

	// fmt.Println(hello.Hello("World", hello.English))

	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(depinj.MyGreeterHandler)))
}
