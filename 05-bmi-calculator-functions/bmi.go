package main

import (
	"fmt"

	"github.com/matheusjerico/bmi/info"
)

func calculateBMI(weigth float32, height float32) (bmiCalculator float32) {
	bmiCalculator = (weigth / (height * height))
	return bmiCalculator
}

func main() {
	info.PrintWelcome()

	weigth, height := getUserMetrics()
	fmt.Printf("Weigth: %.2f, Height: %.2f\n", weigth, height)

	bmiCalculator := calculateBMI(float32(weigth), float32(height))
	printBMI(bmiCalculator)
}
