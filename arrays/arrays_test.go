package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("make the sums of some integer slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tail of integer slices", func(t *testing.T) {
		checkSums(t, SumAllTails([]int{1, 2}, []int{0, 9}), []int{2, 9})
	})
	t.Run("make the sums of tail of integer slices with empty slice", func(t *testing.T) {
		checkSums(t, SumAllTails([]int{}, []int{1, 9, 1}), []int{0, 10})
	})
}

func ExampleSumAll() {
	sums := SumAll([]int{1, 2}, []int{0, 9})
	fmt.Println(sums)
	// Output: [3 9]
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2}, []int{0, 9})
	}
}
