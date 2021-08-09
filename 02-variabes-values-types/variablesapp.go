package main

import (
	"fmt"

	"github.com/matheusjerico/firstapp/greeting"
)

func main() {
	// var greetingText string
	// greetingText = "Hello World!"

	// var greetingText string = "Hello World!"
	// var greetingText = "Hello World!"

	luckyNumber := 17
	evenMoreLuckyNumber := luckyNumber * 20

	fmt.Println(greeting.GreetingText)
	fmt.Println(greeting.GreetingText)
	fmt.Println(luckyNumber)
	fmt.Println(evenMoreLuckyNumber)

	evenMoreLuckyNumber = evenMoreLuckyNumber * 100
	fmt.Println(evenMoreLuckyNumber)

	var newNumber float64 = float64(evenMoreLuckyNumber) / 59
	fmt.Println(newNumber)

	// float 32 vs float64
	var defaultFloat float64 = 1.123456789123456789
	var smallFloat float32 = 1.123456789123456789
	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)

	// default values
	var age int      // declaring the variable
	fmt.Println(age) // outputs 0

	// rune
	var firstRune rune = '$'
	fmt.Println(firstRune)
	fmt.Println(string(firstRune))

	// byte
	var firstByte byte = 'a'
	fmt.Println(firstByte)
	fmt.Println(string(firstByte))

	// string
	firstName := "Matheus"
	lastName := "Palhares"

	fmt.Println(firstName + " " + lastName)
	// fmt.Println(int("9") + 1) CRASH

	fullName := fmt.Sprintf("%s %s", firstName, lastName)
	myAge := 26
	fmt.Printf("Hi, my name is %v (Type: %T) and I am %v (Type: %T) years old.", fullName, fullName, myAge, myAge)
}
