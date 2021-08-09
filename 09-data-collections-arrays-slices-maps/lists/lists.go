package lists

import "fmt"

func lists() {
	var productNames []string
	prices := []float64{
		10.99,
		11.98,
		12.00,
		13.92,
	}

	// productNames
	fmt.Println(productNames)
	productNames = append(productNames, "matheus")
	fmt.Println(productNames)

	// prices
	fmt.Println(prices)
	fmt.Println(prices[1])
	prices[1] = 17
	fmt.Println(prices[1])

	// slices
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
	featuredPrices[0] = 111
	fmt.Println(featuredPrices)
	fmt.Println(prices)

	featuredPricesLast := prices[len(prices)-1]
	fmt.Println(featuredPricesLast)

	// adding new elements list in a list
	discountPrices := []float64{123.2, 234.3, 456.3}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}
