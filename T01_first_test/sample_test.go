// 2020/10/09

package main

import (
	"fmt"
	"testing"
)

func TestSampleFunction(t *testing.T) {
	result := sampleFunction()

	if result == "Success" {
		fmt.Println("Success! Result: ", result)
	} else {
		fmt.Println("Failed... Result: ", result)
	}
}

// End
