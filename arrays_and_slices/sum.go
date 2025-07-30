package arraysandslices

func Sum(arr []int) (sum int) {
	for _, d := range arr {
		sum += d
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, 0, len(numbersToSum))
	for _, s := range numbersToSum {
		sums = append(sums, Sum(s))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	for i, s := range numbersToSum {
		tail := make([]int, 0)
		if len(s) > 0 {
			tail = s[1:]
		}
		sums[i] = Sum(tail)
	}
	return sums
}
