package interaction

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

type RoundData struct {
	Action           string
	CurrentRound     int
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func (roundData *RoundData) PrintRoundStatistics() {
	fmt.Printf("Round: %v\n", roundData.CurrentRound)
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a strong attack against monster for %v damage.\n", roundData.PlayerAttackDmg)

	} else {
		fmt.Printf("Player healed  for %v.\n", roundData.PlayerHealValue)
	}
	fmt.Printf("Monster attacked player for %v damage.\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monster health: %v\n", roundData.MonsterHealth)
}

func NewRoundData(
	action string,
	currentRound int,
	playerAttackDmg int,
	playerHealValue int,
	monsterAttackDmg int,
	playerHealth int,
	monsterHealth int) *RoundData {
	roundData := RoundData{
		Action:           action,
		CurrentRound:     currentRound,
		PlayerAttackDmg:  playerAttackDmg,
		PlayerHealValue:  playerHealValue,
		MonsterAttackDmg: monsterAttackDmg,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
	}
	return &roundData
}

func PrintGreeting() {
	figure := figure.NewFigure("MONSTER PLAYER", "", true)
	figure.Print()
	fmt.Println("Staring a new game...")
	fmt.Println("Good luck!")
}

func ShowAvailableActions(isSpecialRound bool, currentRound int) {
	fmt.Printf("Please choose your action (Round: %v)\n", currentRound)
	fmt.Println("-------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal yourself")
	if isSpecialRound {
		fmt.Println("(3) Special Attack")
	}

}

func DeclarateWinner(winner string) {
	fmt.Println("-------------------------")
	figure := figure.NewColorFigure("GAME OVER", "", "red", true)
	figure.Print()
	fmt.Println("-------------------------")
	fmt.Printf("%v won!\n", winner)
}

func WriteLogFile(rounds *[]RoundData) {

	exPath, err := os.Executable()

	if err != nil {
		fmt.Println("Writing a log file failed. Exiting!")
		return
	}

	exPath = filepath.Dir(exPath)

	file, err := os.Create(exPath + "/gamelog.txt")

	if err != nil {
		fmt.Println("Openning a log file failed. Exiting!")
		return
	}

	for index, value := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                value.Action,
			"Player Attack Damage":  fmt.Sprint(value.PlayerAttackDmg),
			"Player Heal Damage":    fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackDmg),
			"Player Health":         fmt.Sprint(value.PlayerHealth),
			"Monster Health":        fmt.Sprint(value.MonsterHealth),
		}
		logLine := fmt.Sprintln(logEntry)
		_, err = file.WriteString(logLine)

		if err != nil {
			fmt.Println("Writing into log file failed. Exiting!")
			continue
		}
	}

	file.Close()
	fmt.Println("Wrote data to log!")
}
