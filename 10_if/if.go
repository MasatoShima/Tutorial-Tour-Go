// 2020/10/09

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	} else {
		return fmt.Sprint(math.Sqrt(x))
	}
}

// End
