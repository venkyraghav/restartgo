package main

import (
	"fmt"

	hello "github.com/venkyraghav/restartgo/hello"
)

func main() {
	fmt.Println(hello.Hello("World", hello.English))
}
