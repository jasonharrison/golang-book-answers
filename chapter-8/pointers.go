// Write a program that can swap two integers.

package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println("a =", a, "b =", b)
}