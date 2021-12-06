package main

import (
	"fmt"
	"strconv"
)

func getPermutation(n int, k int) string {
	used := make([]bool, n+1)
	result := ""
	var backtrack func(total int, index int)
	backtrack = func(total int, index int) {
		if index == n {
			return
		}
		count := factorial(n - index - 1)
		for i := 1; i <= n; i++ {
			if used[i] {
				continue
			}
			if count < total {
				total -= count
				continue
			}
			used[i] = true
			result = result + strconv.Itoa(i)
			backtrack(total, index+1)
			return
		}
	}
	backtrack(k, 0)
	return result
}

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(getPermutation(4, 9))
}
