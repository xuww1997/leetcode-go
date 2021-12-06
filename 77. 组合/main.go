package main

import "fmt"

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	var backtrack func(begin int)
	tmp := make([]int, 0)
	backtrack = func(begin int) {
		if len(tmp) == k {
			r := make([]int, len(tmp))
			copy(r, tmp)
			result = append(result, r)
		} else {
			for i := begin; i <= n; i++ {
				tmp = append(tmp, i)
				backtrack(i + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtrack(1)
	return result
}

func main() {
	fmt.Println(combine(1, 1))
}
