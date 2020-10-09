// 2020/10/09

package main

import (
	"golang.org/x/tour/pic"
)

func main() {
	pic.Show(Pic)
}

func Pic(dx int, dy int) [][]uint8 {
	p := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		p[y] = make([]uint8, dx)

		for x := 0; x < dx; x++ {
			p[y][x] = uint8((x ^ y) * (x + y) / 2)
		}
	}

	return p
}

// End
