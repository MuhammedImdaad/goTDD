package arrays

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [...]int{1, 2, 3, 4, 5} // An array's length is part of its type, so arrays cannot be resized.

		got := Sum(numbers[:])
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3} //slice points to the [3]int underlying array

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if slices.Equal(got, want) == false {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSlicing(t *testing.T) {
	// Create a slice with 6 elements
	s := []int{2, 3, 5, 7, 11, 13}

	// Slice the slice to give it zero length (s[0:0], start=0, end=0)
	s = s[:0]
	want := []int{}
	printSlice(t, s, want)

	// Extend its length to 4 (s[0:4], start=0, end=4)
	s = s[:4]
	want = []int{2, 3, 5, 7}
	printSlice(t, s, want)

	// Drop its first two values (s[2:], start=2, end=4)
	s = s[2:]
	want = []int{5, 7}
	printSlice(t, s, want)

	// Take the next element (s[2:3], start=2, end=3)
	s = s[2:3]
	want = []int{11}
	printSlice(t, s, want)

	// Extend to include the next element (s[:2], start=0, end=2)
	s = s[:2]
	want = []int{11, 13}
	printSlice(t, s, want)
}

func printSlice(t testing.TB, got, want []int) {
	t.Helper()
	if slices.Equal(got, want) == false {
		t.Errorf("got %v want %v", got, want)
	}
	fmt.Printf("len=%d cap=%d %v\n", len(got), cap(got), got)
}
