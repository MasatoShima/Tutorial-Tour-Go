// 2020/11/15

package main

import (
	"strings"
)

import (
	"golang.org/x/tour/wc"
)

func main() {
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	Counter := make(map[string]int)

	words := strings.Split(s, " ")

	for _, word := range words {
		_, ok := Counter[word]

		if ok == true {
			Counter[word]++
		} else {
			Counter[word] = 1
		}
	}

	return Counter
}

// End
