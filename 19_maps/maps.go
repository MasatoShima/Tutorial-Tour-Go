// 2020/11/12

package main

import "fmt"

type Vertex struct {
	Lat  float64
	Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		Lat:  5.0,
		Long: 8.0,
	}

	fmt.Println(m["Bell Labs"])
}

// End
