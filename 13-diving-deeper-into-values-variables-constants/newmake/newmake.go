package newmake

import "fmt"

func main() {
	// new & make
	fmt.Println("------------------------")
	fmt.Println("new")
	fmt.Println("------------------------")
	number := new(int)
	fmt.Println(number)
	fmt.Println(*number)

	anotherNumber := 0
	addressNumber := &anotherNumber
	fmt.Println(addressNumber)
	fmt.Println(anotherNumber)

	moreHobbies := new([]string)

	fmt.Println(moreHobbies)
	fmt.Println(*moreHobbies)
	*moreHobbies = append(*moreHobbies, "Cooking", "Dancing")
	fmt.Println(*moreHobbies)

	fmt.Println("------------------------")
	fmt.Println("make")
	fmt.Println("------------------------")
	// hobbies := []string{}
	// hobbies := make([]string,)
	hobbies := make([]string, 2, 10)

	// aMap := make(map[string]int, 5)

	fmt.Println(hobbies)
	hobbies[0] = "Sports"
	hobbies[1] = "Reading"
	hobbies = append(hobbies, "Cooking", "Dancing")

	fmt.Println(hobbies)
}
