package main

import "fmt"

func maxProfit(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if (price - minPrice) > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
