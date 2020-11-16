// 2020/11/15

package main

import "fmt"

type T struct {
	S string
}

type I interface {
	M()
}

func main() {
	var i I = T{"Hello"}
	i.M()
}

func (t T) M() {
	fmt.Println(t.S)
}

// End
