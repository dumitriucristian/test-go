package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {

	lengthOfNum := len(numbersToSum)
	sums := make([]int, lengthOfNum)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
