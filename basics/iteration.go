package basics

import "strings"

func Repeat(input string, repeat int) string {
	// var output string //Strings in Go are immutable,
	// for i := 0; i < repeat; i++ {
	// 	output += input
	// }
	// return output

	var output strings.Builder
	for range repeat {
		output.WriteString(input)
	}
	
	return output.String()
}
