package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/matheusjerico/bmi/info"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics() (weigth float64, height float64) {
	weigth = getUserInput(info.WeigthPrompt, 32)
	height = getUserInput(info.HeightPrompt, 32)
	return weigth, height
}

func getUserInput(promptText string, bitSize int) (value float64) {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	userInput = strings.Replace(userInput, ",", ".", -1)
	value, _ = strconv.ParseFloat(userInput, bitSize)

	return value
}
