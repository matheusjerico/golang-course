package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoise(isSpecialRound bool, currentRound int) string {
	for {
		playerChoise, _ := getPlayerInput()

		if playerChoise == "1" {
			return "ATTACK"
		} else if playerChoise == "2" {
			return "HEAL"
		} else if playerChoise == "3" && isSpecialRound {
			return "SPECIAL_ATTACK"
		} else {
			fmt.Println("WARNING!!! Fetching the user input failed. Please try again")
			ShowAvailableActions(isSpecialRound, currentRound)
		}
	}

}

func getPlayerInput() (string, error) {
	fmt.Print("Your choise: ")
	playerInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	playerInput = strings.Replace(playerInput, "\n", "", -1)

	return playerInput, nil
}
