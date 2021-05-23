package slices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	sums := make([]int, len(numbers))

	for i, n := range numbers {
		sums[i] = Sum(n)
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	sums := make([]int, len(numbers))

	for i, n := range numbers {
		if len(n) == 0 {
			sums[i] = 0
		} else {
			sums[i] = Sum(n[1:])
		}
	}

	return sums
}
