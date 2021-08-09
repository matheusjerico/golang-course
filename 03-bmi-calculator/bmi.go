package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matheusjerico/bmi/info"
)

func replaceValues(weigth string, height string) (weigthReplaced string, heightReplaced string) {
	weigthReplaced = strings.Replace(weigth, "\n", "", -1)
	weigthReplaced = strings.Replace(weigthReplaced, ",", ".", -1)
	heightReplaced = strings.Replace(height, "\n", "", -1)
	heightReplaced = strings.Replace(heightReplaced, ",", ".", -1)
	return weigthReplaced, heightReplaced
}

func convertStringFloat(weigth string, height string, bitSize int) (weigthNumber float64, heightNumber float64) {
	weigthNumber, _ = strconv.ParseFloat(weigth, bitSize)
	heightNumber, _ = strconv.ParseFloat(height, bitSize)
	return weigthNumber, heightNumber
}

func main() {
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)
	fmt.Print(info.WeigthPrompt)
	weigthInput, _ := reader.ReadString('\n')
	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	weigthInput, heightInput = replaceValues(
		weigthInput, heightInput)

	weigthNumber, heightNumber := convertStringFloat(
		weigthInput, heightInput, 32)

	fmt.Printf("Weigth: %.2f, Height: %.2f\n", weigthNumber, heightNumber)

	bmiCalculator := (weigthNumber / (heightNumber * heightNumber))
	fmt.Printf("Your BMI: %.2f\n", bmiCalculator)
}
