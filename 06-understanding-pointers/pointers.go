package main

import "fmt"

func main() {
	age := 32
	fmt.Println(age)
	fmt.Println(&age)
	fmt.Println("-------")

	// myAge := &age
	// var myAge *int
	myAge := &age
	fmt.Println(myAge)
	fmt.Println(*myAge)
	fmt.Println("-------")
	*myAge = 27
	fmt.Println(myAge)
	fmt.Println(*myAge)
	fmt.Println(&age)
	fmt.Println(age)

	fmt.Println("-------")
	fmt.Println(double(myAge))
	fmt.Println(*myAge)
}

func double(number *int) int {
	result := *number * 2
	*number = 100
	return result
}
