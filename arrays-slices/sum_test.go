package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Slice of arbitrary length", func(t *testing.T) {
		// 'slice' has arbitrary length
		// note the syntax with no size specified
		// []int and [5]int and [4]int are all different types!
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}
