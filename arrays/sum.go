package array

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(xss ...[]int) []int {
	var sum []int
	for _, xs := range xss {
		sum = append(sum, Sum(xs))
	}
	return sum
}

func SumTail(xss ...[]int) []int {
	var sum []int
	for _, xs := range xss {
		if len(xs) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(xs[1:]))
		}
	}
	return sum
}
