package main

import (
	"reflect"
	"slices"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2}
	slice2 := []int{0, 9}
	got := SumAll(slice1, slice2)
	want := []int{3, 9}
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	compareSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("sum slices", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{0, 9}
		got := SumAllTails(slice1, slice2)
		want := []int{2, 9}
		compareSums(t, got, want)
	})
	t.Run("empty slice returns {0}", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{0, 9}
		got := SumAllTails(slice1, slice2)
		want := []int{0, 9}
		compareSums(t, got, want)
	})
}
