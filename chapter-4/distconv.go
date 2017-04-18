// Write another program that converts from feet into meters.

package main

import "fmt"

func ConvertToMeters(ft float64) float64 {
	var m = ft*0.3048
	return m
}

func main() {
	fmt.Print("Enter a distance in feet: ")
	var input float64
	fmt.Scanf("%f", &input)

	var output float64 = ConvertToMeters(input)
	fmt.Println(input, "ft = ", output, "m")
}
