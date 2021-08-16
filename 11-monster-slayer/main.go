package main

import (
	"fmt"

	"github.com/matheusjerico/monsterslayer/actions"
	"github.com/matheusjerico/monsterslayer/interaction"
)

var currentRound = 0
var gameRounds = []interaction.RoundData{}

func main() {
	startGame()

	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
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

	var playerAttackDmg int
	var playerHealValue int
	var monsterAttackValue int

	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		playerAttackDmg = actions.AttackMonster(true)
	}

	monsterAttackValue = actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthAmounts()

	roundData := interaction.NewRoundData(
		userChoice,
		currentRound,
		playerAttackDmg,
		playerHealValue,
		monsterAttackValue,
		playerHealth,
		monsterHealth,
	)
	roundData.PrintRoundStatistics()
	gameRounds = append(gameRounds, *roundData)

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclarateWinner(winner)
	interaction.WriteLogFile(&gameRounds)
}
