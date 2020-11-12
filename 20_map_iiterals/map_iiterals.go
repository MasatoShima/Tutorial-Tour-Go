// 2020/11/12

package main

import "fmt"

type Vertex struct {
	Lat  float64
	Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {
		Lat:  5.0,
		Long: 8.0,
	},
	"Google": {
		Lat:  50.0,
		Long: 80.0,
	},
}

func main() {
	fmt.Println(m)
}

// End
