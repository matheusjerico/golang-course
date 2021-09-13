package variadic

import "fmt"

func main() {
	numbers := []int{1, 10, 15}

	sum := sumUp(1, 10, 15)
	anotherSum := sumUp(1, numbers...)

	fmt.Printf("Sum up: %v\n", sum)
	fmt.Printf("Sum up: %v\n", anotherSum)
}

func sumUp(startingValue int, numbers ...int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	return sum
}
