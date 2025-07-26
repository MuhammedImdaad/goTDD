package arrays

// a generic reduce funtion with a collection and a reducer
func Reduce[T any](collection []T, f func(T, T) T, initialValue T) T {
	var result T = initialValue
	for _, v := range collection {
		result = f(result, v)
	}

	return result
}

func Sum(numbers []int) int { // accepts a slice

	// Version II
	// Define a reducer function that adds two integers
	f := func(acc, in int) int { return acc + in }
	// Use the generic Reduce function to sum all elements in the slice
	return Reduce(numbers, f, 0)
}

// variadic function that can take a variable number of slices
func SumAll(nums ...[]int) []int {

	// Version II
	// Define a reducer function that appends the sum of each slice to the accumulator
	f := func(acc, in []int) []int {
		return append(acc, Sum(in))
	}
	// Use the generic Reduce function to process all input slices and collect their sums
	return Reduce(nums, f, []int{})
}
