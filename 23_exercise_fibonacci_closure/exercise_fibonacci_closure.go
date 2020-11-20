// 2020/11/15

package main

import "fmt"

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	n, m := 0, 1
	return func() int {
		i := n
		n, m = m, n+m
		return i
	}
}

// End
