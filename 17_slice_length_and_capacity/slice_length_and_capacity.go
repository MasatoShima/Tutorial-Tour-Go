// 2020/10/09

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:5]
	printSlice(s)

	// Drop its first two values.
	s = s[3:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf(
		"length: %d capacity: %d %v \n",
		len(s),
		cap(s),
		s,
	)
}

// End
