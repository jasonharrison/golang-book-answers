// Write a function with one variadic parameter that finds the greatest number in a list of numbers.

package main

import (
	"fmt"
)

func greatest(args ...int) int {
	var g int = 0
	for _, x := range args {
		if x > g {
			g = x
		}
	}
	return g
}

func main() {
	fmt.Println(greatest(2, 0, 1, 7))
}
