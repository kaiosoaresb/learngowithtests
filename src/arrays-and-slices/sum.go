package arrays_and_slices

func Sum(numbers []int) int {
	var sum int
	//for i := 0; i < len(numbers); i++ {
	//	sum += numbers[i]
	//}

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll1(numbersToSum ...[]int) []int {

	lenghtOfNumbers := len(numbersToSum)
	sums := make([]int, lenghtOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func SumAll2(numbersToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {

	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
