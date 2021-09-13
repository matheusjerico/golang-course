package anonymous

import "fmt"

type transformFn func(int) int

// type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	transformed := transformNumber(&numbers, func(number int) int {
		return number * 2
	})
	fmt.Println(transformed)

	double := createTransformer(2)
	triple := createTransformer(3)

	doubled := transformNumber(&numbers, double)
	tripled := transformNumber(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumber(numbers *[]int, transform transformFn) []int {
	dNumbers := make([]int, 0)

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
