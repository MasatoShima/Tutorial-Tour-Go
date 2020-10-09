// 2020/10/09

package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("Hello")
}

// End
