// Write a function which takes an integer and halves it and returns true if it was even or false if it was odd.

package main

import "fmt"

func halve(x int) bool {
	x = x/2
	if x % 2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanf("%d", &input)
	result := halve(input)
	if result {
		fmt.Println(input/2, "is even")
	} else {
		fmt.Println(input/2, "is odd")
	}
}