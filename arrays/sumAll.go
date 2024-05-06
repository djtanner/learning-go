package main

// variadic function takes a vaariable number of args
// sums each slice and returns a slice of the sums
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums

}
