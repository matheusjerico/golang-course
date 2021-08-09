package main

import "fmt"

func main() {
	// 2)
	var piNumber float64 = 3.14
	var circleRadius float64 = 5

	// 3)
	var circleCircumference float64 = 2 * piNumber * circleRadius
	fmt.Printf("Circule circumference: %v (Type: %T)\n", circleCircumference, circleCircumference)

	//  4)
	fmt.Printf("For a radius of %.0f, the circle circumference is %.2f.\n", circleRadius, circleCircumference)
}
