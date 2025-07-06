package arrays

func Sum(numbers []int) int { // accepts a slice
	var result int
	for _, v := range numbers { // ignore the index value by using _
		result += v
	}

	return result
}

// variadic functions that can take a variable number of arguments
func SumAll(nums ...[]int) []int {
	// len := len(nums)
	// sums := make([]int, len) //make function; this is how you create dynamically-sized arrays.
	// //The make function allocates a zeroed array and returns a slice that refers to that array

	// for i, v := range nums {
	// 	sums[i] = Sum(v)
	// }

	// return sums

	var sums []int
	for _, v := range nums {
		sums = append(sums, Sum(v))
	}

	return sums
}
