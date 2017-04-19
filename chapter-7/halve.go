// Write a function which takes an integer and halves it and returns true if it was even or false if it was odd.

package main

import "fmt"

func halve(x int) (int, bool) {
	x = x/2
	if x % 2 == 0 {
		return x, true
	} else {
		return x, false
	}
}

func main() {
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanf("%d", &input)
	halved, result := halve(input)
	if result {
		fmt.Println(halved, "is even")
	} else {
		fmt.Println(halved, "is odd")
	}
}