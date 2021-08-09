package info

import "fmt"

const (
	mainTitle    = "BMI Calculator!"
	separator    = "--------------------"
	WeigthPrompt = "Please enter your weight (kg): "
	HeightPrompt = "Please enter your height (m): "
)

func PrintWelcome() {
	fmt.Println(mainTitle)
	fmt.Println(separator)
}
