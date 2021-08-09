package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a, b := generateRandomNumbers(110)
	sum := add(a, b)
	printNumber(sum)
}

func add(a int, b int) (sum int) {
	sum = a + b
	return
}

func generateRandomNumbers(n int) (randomNumberOne int, randomNumberTwo int) {
	randomNumberOne = rand.Intn(n)
	randomNumberTwo = rand.Intn(n)
	return
}

func printNumber(number int) {
	fmt.Printf("The number is %v", number)
}
