package main

import (
	"fmt"

	"github.com/matheusjerico/monsterslayer/interaction"
)

var currentRound = 0

func main() {
	startGame()

	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound, currentRound)
	userChoice := interaction.GetPlayerChoise(isSpecialRound, currentRound)

	fmt.Printf("User choice: %v\n", userChoice)

	return ""
}

func endGame() {

}
