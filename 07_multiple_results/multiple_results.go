// 2020/10/08

package main

import "fmt"

func main() {
	a, b := swap("x", "y")
	fmt.Printf("%s %s", a, b)
}

func swap(x string, y string) (string, string) {
	return y, x
}

// End
