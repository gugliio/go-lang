package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var suma []int
	for _, n := range numbersToSum {
		suma = append(suma, Sum(n))
	}
	return suma
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sumTail []int
	for _, num := range numbersToSum {
		if len(num) == 0 {
			sumTail = append(sumTail, 0)
		} else {
			tail := num[1:]
			sumTail = append(sumTail, Sum(tail))
		}
	}
	return sumTail
}
