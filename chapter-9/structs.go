// Demonstrate structs and methods

package main

import "fmt"
import "math"

type circle struct {
	x, y, r float64
}

func (c *circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := circle{x: 0, y: 0, r: 5}
	fmt.Println(c.area())
}
