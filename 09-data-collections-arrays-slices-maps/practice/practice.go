package main

import "fmt"

// 6
type Product struct {
	title string
	id    string
	price float64
}

func NewProduct(title string, id string, price float64) Product {
	product := Product{
		title: title,
		id:    id,
		price: price,
	}
	return product
}

func main() {
	// 1)
	hobbies := []string{}
	hobbies = append(hobbies, "futebol", "tenis", "volei", "corrida")
	fmt.Println(hobbies)

	// 2)
	firstElement := hobbies[0]
	fmt.Printf("First element: %v\n", firstElement)
	secondThirdElements := hobbies[1:3]
	fmt.Printf("Second and third elements: %v\n", secondThirdElements)

	// 3)
	firstSecondElements := hobbies[0:2]
	fmt.Printf("First and second: %v\n", firstSecondElements)
	firstSecondElements = hobbies[:2]
	fmt.Printf("First and second: %v\n", firstSecondElements)

	// 4)
	firstSecondElements = append(firstSecondElements, hobbies[2])
	fmt.Printf("%v\n", firstSecondElements)

	// 5)
	courseGoals := []string{
		"learn",
		"goland",
	}
	fmt.Println(courseGoals)
	courseGoals[1] = "gitlab"
	fmt.Println(courseGoals)
	courseGoals = append(courseGoals, "airflow")
	fmt.Println(courseGoals)

	// 6)
	ProductsList := []Product{
		NewProduct(
			"maquina de lavar",
			"1",
			1234.2,
		),
		NewProduct(
			"robo aspirador",
			"2",
			2345.6,
		),
	}
	fmt.Println(ProductsList)

	// 7)
	ProductsList = append(ProductsList, NewProduct(
		"ps5",
		"3",
		5432.2,
	))
	fmt.Println(ProductsList)

}
