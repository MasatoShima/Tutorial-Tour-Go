// 2020/10/09

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)

	x := v.X
	fmt.Println(x)
}

// End
