package main

import (
	"fmt"
)

func main() {
	fmt.Println(Reverse("cool"))
}

// BenchmarkReverse-8   	33887774	        36.91 ns/op	       0 B/op	       0 allocs/op
// BenchmarkReverse-8   	 9968499	       126.3 ns/op	      12 B/op	       4 allocs/op
func Reverse(input string) string {
	// convert string to rune slice
	var runes = []rune(input)

	// reverse the rune slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// convert rune slice back to string
	return string(runes)
}
