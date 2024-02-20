package main

import (
	"os"

	"github.com/venkyraghav/restartgo/mock"
)

func main() {
	mock.Countdown(os.Stdout)
	// fmt.Println(hello.Hello("World", hello.English))

	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(depinj.MyGreeterHandler)))
}
