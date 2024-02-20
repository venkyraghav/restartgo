package mock

import (
	"fmt"
	"io"
)

func Countdown(writer io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintf(writer, "%d\n", i)
	}
}
