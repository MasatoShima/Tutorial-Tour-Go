// 2020/10/09

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v1 := Vertex{X: 1, Y: 2}
	fmt.Println(v1)

	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{}
	fmt.Println(v3)

	v4 := &Vertex{X: 1, Y: 2}
	fmt.Println(v4)
}

// End
