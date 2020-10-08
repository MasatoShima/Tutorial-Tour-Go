// 2020/10/09

package main

import "fmt"

func main() {
	for i := 0.0; i < 10.0; i++ {
		result := sqrt(i)
		fmt.Println(result)
	}
}

func sqrt(x float64) float64 {
	z := 1.0
	z -= (z*z - x) / (2 * z)
	return z
}

// End
