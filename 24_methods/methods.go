// 2020/11/15

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

func main() {
	value := Vertex{5, 8}
	fmt.Println(value.Abs())
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// End
