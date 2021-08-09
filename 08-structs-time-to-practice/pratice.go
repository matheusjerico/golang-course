package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// 1) Create a new type / data structure for storing product data (e.g. about a book)
type Book struct {
	Id          int
	TitleName   string
	Description string
	Price       float64
}

func (book *Book) outputBookDetails() {
	fmt.Printf("The ID is: %v,\nThe book name is: %v\nThe Description is: %v,\nThe Price is: %v\n",
		book.Id,
		book.TitleName,
		book.Description,
		book.Price)
}

func (book *Book) writeJson() {
	fileName := fmt.Sprintf("%v.json", book.Id)

	file, _ := json.MarshalIndent(book, "", " ")
	_ = ioutil.WriteFile(fileName, file, 0644)
}

func (book *Book) writeText() {
	file, _ := os.Create(fmt.Sprint("$v.txt", book.Id))
	content := fmt.Sprintf("The ID is: %v,\nThe book name is: %v\nThe Description is: %v,\nThe Price is: %v\n",
		book.Id,
		book.TitleName,
		book.Description,
		book.Price)

	file.WriteString(content)
	file.Close()
}

func NewBook(Id int, tName string, Description string, Price float64) *Book {
	book := Book{
		Id:          Id,
		TitleName:   tName,
		Description: Description,
		Price:       Price,
	}

	return &book
}

func main() {
	// 2) Create concrete instances of this data type in two different ways:
	myBook := Book{
		Id:          1,
		TitleName:   "Harry Potter",
		Description: "Harry Potter is a 2000's movies that all people love",
		Price:       10,
	}
	fmt.Println(myBook)

	var myBookPointer *Book
	myBookPointer = NewBook(
		2,
		"Harry Potter 2",
		"Harry Potter is a 2000's movies that all people love",
		11,
	)
	fmt.Println(*myBookPointer)

	// 3) Add a "connected function" that outputs the data + call that function from insIde "main"
	myBookPointer.outputBookDetails()

	// 4) Change the program to fetch user input values for the different data fields
	book := getBook()
	book.outputBookDetails()

	// 5) Bonus: Add a connected "store" function that writes that data into a file
	book.writeJson()
	book.writeText()
}

func getBook() *Book {
	bookId := getUserData("Please enter your book Id: ")
	bookTitleName := getUserData("Please enter your book title name: ")
	bookDescription := getUserData("Please enter your book Description: ")
	bookPrice := getUserData("Please enter your book Price: ")
	bookIdInt, _ := strconv.Atoi(bookId)
	bookPriceFloat, _ := strconv.ParseFloat(bookPrice, 64)
	book := NewBook(
		bookIdInt,
		bookTitleName,
		bookDescription,
		bookPriceFloat,
	)

	return book
}

func getUserData(promptText string) (cleanedInput string) {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	cleanedInput = strings.Replace(userInput, "\n", "", -1)

	return cleanedInput
}
