package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	min := math.MaxInt32
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}
	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
