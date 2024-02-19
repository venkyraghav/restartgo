package integers

import (
	"fmt"
	"testing"
)

func assertCorrectAddition(t *testing.T, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
func TestAdd(t *testing.T) {
	t.Run("2 + 2", func(t *testing.T) {
		assertCorrectAddition(t, Add(2, 2), 4)
	})
	t.Run("3 + 3", func(t *testing.T) {
		assertCorrectAddition(t, Add(3, 3), 6)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
