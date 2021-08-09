package main

import "fmt"

func main() {
	// var firstName string = "Matheus"
	// var lastName string = "Palhares"
	firstName := "Matheus"
	lastName := "Palhares"

	fmt.Printf("My name is %s %s\n", firstName, lastName)

	// var currentYear int = 2021
	// var birthYear int = 1994
	currentYear := 2021
	birthYear := 1994
	differenceYear := currentYear - birthYear

	fmt.Printf("Current year: %d, birth year: %d\n", currentYear, birthYear)
	fmt.Printf("The difference between current and birth year is %d\n", differenceYear)

	currentYear += 1
	differenceYear = currentYear - birthYear
	fmt.Printf("Current year: %d, birth year: %d\n", currentYear, birthYear)
	fmt.Printf("The difference between current and birth year is %d\n", differenceYear)

}
