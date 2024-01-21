package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {

	numbers1 := []int{1, 2}
	numbers2 := []int{0, 9}

	got := SumAll(numbers1, numbers2)
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
