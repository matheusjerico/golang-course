package main

import "fmt"

func main() {
	websites := map[string]string{
		"google":              "https://google.com",
		"amazon web services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["google"])
	fmt.Println(websites["amazon web services"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "amazon web services")
	fmt.Println(websites)
}
