package functionsarevalues

import "fmt"

type transformFn func(int) int

// type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{5, 3, 1}
	doubled := transformNumber(&numbers, double)
	tripled := transformNumber(&numbers, triple)
	fmt.Println(numbers)
	fmt.Println(doubled)
	fmt.Println(tripled)
	fmt.Println("============================================")

	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&moreNumbers)

	transformedNumber1 := transformNumber(&numbers, transformFn1)
	transformedNumber2 := transformNumber(&moreNumbers, transformFn2)
	fmt.Println(transformedNumber1)
	fmt.Println(transformedNumber2)
}

func transformNumber(numbers *[]int, transform transformFn) []int {
	dNumbers := make([]int, 0)

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
