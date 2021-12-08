package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > buy {
			profit += prices[i] - buy
		}
		buy = prices[i]
	}
	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
