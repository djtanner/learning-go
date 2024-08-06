// Write your answer here, and then test your code.
// Your job is to implement the CalculateMean() function.

package main

import (
	"errors"
	"fmt"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

// CalculateMean returns mean of the numbers slice, or an error.
func CalculateMean[T float64 | int](numbers []T) (*float64, error) {
    // Your code goes here.

    if len(numbers) == 0 {
        return nil, errors.New("invalid input")
    }

    var result T

    for i :=0; i < len(numbers); i++{
        result += numbers[i]
    }

    mean := float64(result) / float64(len(numbers))
    return &mean, nil
}

func main(){
	numbers := []int{1, 2, 3, 4, 5}
	result, err := CalculateMean(numbers)
	fmt.Println(*result)
	if err != nil {
		panic(err)
	}
	if showExpectedResult {
		println(*result) // 3
	}
	if showHints {
		println("The mean of 1, 2, 3, 4, 5 is 3.")
	}
}
