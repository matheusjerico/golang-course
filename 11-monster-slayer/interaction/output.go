package interaction

import "fmt"

func PrintGreeting() {
	fmt.Println("MONSTER PLAYER")
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
