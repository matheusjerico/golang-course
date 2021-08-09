package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	selectedChoice, err := getUserChoice()

	if err != nil {
		fmt.Println("Invalid choice, exiting!")
		return
	}

	switch selectedChoice {
	case "1":
		calculateSumUpToNumber()
	case "2":
		calculateFactorial()
	case "3":
		calculateSumManually()
	default:
		calculateListSum()
	}
}

func getInputNumber() (int, error) {
	inputNumber, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Invalid number")
		return 0, err
	}

	inputNumber = strings.Replace(inputNumber, "\n", "", -1)
	chosenNumber, err := strconv.ParseInt(inputNumber, 0, 64)

	if err != nil {
		return 0, err
	}

	return int(chosenNumber), nil
}

func getUserChoice() (string, error) {
	fmt.Println("Please make your choice: ")
	fmt.Println("1) Add up all the number of to number X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) Sum up manually entered numbers")
	fmt.Println("4) Sum up a list of entered numbers")
	fmt.Print("Please make your choice: ")
	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	userInput = strings.Replace(userInput, "\n", "", -1)

	if userInput == "1" ||
		userInput == "2" ||
		userInput == "3" ||
		userInput == "4" {
		return userInput, nil
	} else {
		return "", errors.New("INVALID INPUT")
	}
}

func calculateSumUpToNumber() {
	fmt.Print("Please enter your number: ")
	chosenNumber, err := getInputNumber()
	sum := 0

	if err != nil {
		fmt.Println("Invalid Input")
	}

	for i := 1; i <= chosenNumber; i++ {
		sum = sum + i
	}

	fmt.Printf("Result: %d", sum)

}

func calculateFactorial() {
	fmt.Print("Please enter your number: ")
	chosenNumber, err := getInputNumber()
	factorial := 1

	if err != nil {
		fmt.Println("Invalid Input")
	}

	for i := chosenNumber; i > 0; i-- {
		factorial = factorial * i
	}

	fmt.Printf("Result: %d", factorial)
}

func calculateSumManually() {
	isEnteringNumbers := true
	sum := 0

	fmt.Println("Keep entering number, the program will quite once you enter any other value.")
	for isEnteringNumbers {
		chosenNumber, err := getInputNumber()
		isEnteringNumbers = err == nil

		sum = sum + chosenNumber
	}

	fmt.Printf("Result: %d", sum)
}

func calculateListSum() {
	fmt.Println("Please enter a comma separated list of numbers.")
	inputListNumber, err := reader.ReadString('\n')
	sum := 0

	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	inputListNumber = strings.Replace(inputListNumber, "\n", "", -1)
	inputListNumber = strings.Replace(inputListNumber, " ", "", -1)
	inputNumbers := strings.Split(inputListNumber, ",")

	for index, value := range inputNumbers {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
		intNumber, err := strconv.Atoi(value)

		if err != nil {
			continue
		}

		sum = sum + intNumber
	}

	fmt.Printf("Result: %d", sum)
}
