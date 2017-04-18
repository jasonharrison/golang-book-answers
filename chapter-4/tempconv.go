// Write a program that converts from Fahrenheit into Celsius.

package main

import "fmt"

func ConvertToFahrenheit(f float64) float64 {
	var c = (f - 32) * 5/9
	return c
}

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	var output float64 = ConvertToFahrenheit(input)
	fmt.Println(input, "°F = ", output, "°C")
}
