package arrays

// Sum - sums up all integers in input array
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAll - takes in variable number of integer arrays and returns their sums
func SumAll(numArraysToSum ...[]int) []int {
	sums := []int{}

	for _, numbers := range numArraysToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails returns the sums of all elements of an array but the first one
func SumAllTails(numsToSum ...[]int) []int {
	sums := []int{}

	for _, numbers := range numsToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}

	return sums
}
