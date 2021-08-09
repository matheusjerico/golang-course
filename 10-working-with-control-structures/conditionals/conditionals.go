package conditionals

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func conditionals() {

	userAge, err := getUserAge()

	if err != nil {
		fmt.Println("Invalid input! Enter a valid number!")
		return
	}

	// isOldEnough := userAge < 18
	// if isOldEnough {
	// 	fmt.Printf("Welcome to the kids club: %v age\n", userAge)
	// }

	if (userAge >= 30 && userAge < 50) || userAge > 60 {
		fmt.Printf("You are eligible for our senior jobs: %v age\n", userAge)
	} else if userAge >= 50 {
		fmt.Printf("The best age: %v age\n", userAge)
	} else if userAge >= 18 {
		fmt.Printf("Welcome to the club: %v age\n", userAge)
	} else {
		fmt.Printf("Sorry, you are not old enough:/ %v age\n", userAge)
	}

	// // switch case
	// switch userAge {
	// case 1:
	// 	fmt.Println("Option 1")
	// case 2:
	// 	fmt.Println("Option 2")
	// case 3:
	// 	fmt.Println("Option 3")
	// default:
	// 	fmt.Println("Invalid choice!")
	// }
	// fmt.Println("Goodbye!!")
}

func getUserAge() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your age: ")
	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.Replace(userAgeInput, "\n", "", -1)
	userAge, err := strconv.ParseInt(userAgeInput, 0, 64)

	return int(userAge), err
}
