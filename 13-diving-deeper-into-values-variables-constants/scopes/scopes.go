package scopes

import "fmt"

var userName = "Matheus"

func main() {
	shouldContinue := true
	var userName = "Jericó"

	if shouldContinue {
		userAge := 27

		fmt.Printf("Name: %v, Age: %v\n", userName, userAge)
	}
	printData()

}

func printData() {
	fmt.Printf("Name: %v\n", userName)
}
