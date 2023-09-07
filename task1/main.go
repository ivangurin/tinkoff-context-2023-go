package main

import (
	"fmt"
)

func main() {

	var counter, money int

	fmt.Scanln(&counter, &money)

	prices := make([]int, 0, counter)
	price := 0
	for i := 0; i < counter; i++ {
		fmt.Scan(&price)
		prices = append(prices, price)
	}

	bestPrice := 0
	for _, price = range prices {
		if price <= money && price > bestPrice {
			bestPrice = price
		}
	}

	fmt.Println(bestPrice)

}
