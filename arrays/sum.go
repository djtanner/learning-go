package main

func Sum(numbers []int) int {
	sum := 0
	//the blank identifier _ is used to ignore the index
	for _, number := range numbers {
		sum += number
	}

	return sum
}
