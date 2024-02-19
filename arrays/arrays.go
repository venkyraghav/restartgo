package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(arraysToSum ...[]int) []int {
	var sums []int
	for _, array := range arraysToSum {
		sums = append(sums, Sum(array))
	}
	return sums
}

func SumAllTails(arraysToSum ...[]int) []int {
	var sums []int
	for _, array := range arraysToSum {
		if len(array) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(array[1:]))
		}
	}
	return sums
}
