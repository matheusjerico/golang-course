package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

type User struct {
	firstName   string
	lastName    string
	birthdate   string
	createdDate time.Time
}

func (user *User) outputUserDetails() {
	fmt.Printf("My name is is %s %s (born on %s)", user.firstName, user.lastName, user.birthdate)
}

func main() {
	var newUser *User
	// newUser.firstName = getUserData("Please enter your first name: ")
	// newUser.lastName = getUserData("Please enter your last name: ")
	// newUser.birthdate = getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	// newUser.createdDate = time.Now()
	// fmt.Println(newUser)

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	newUser = NewUser(firstName, lastName, birthdate)
	newUser.outputUserDetails()
}

func NewUser(fName string, lName string, bDate string) *User {
	created := time.Now()

	user := User{
		firstName:   fName,
		lastName:    lName,
		birthdate:   bDate,
		createdDate: created,
	}
	return &user
}

func getUserData(promptText string) (cleanedInput string) {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	cleanedInput = strings.Replace(userInput, "\n", "", -1)

	return cleanedInput
}
