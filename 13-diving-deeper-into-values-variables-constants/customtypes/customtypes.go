package customtypes

import "fmt"

type Person struct {
	name string
	age  int
}

type customNumber int

type personData map[string]Person

func (number customNumber) pow(power int) customNumber {
	result := number

	for i := 1; i < power; i++ {
		result = result * number
	}

	return result
}

func main() {
	var people personData = personData{
		"p1": {"Matheus", 27},
	}

	fmt.Println(people)

	var number customNumber = 5
	fmt.Println(number)
	number = number.pow(3)
	fmt.Println(number)

}
