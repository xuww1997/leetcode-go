package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0)
	tmp := make([]int, 0)
	var backtrack func(begin int, t int)
	backtrack = func(begin int, t int) {
		if t < 0 {
			return
		} else if t == 0 {
			r := make([]int, len(tmp))
			copy(r, tmp)
			result = append(result, r)
		} else {
			for i := begin; i < len(candidates); i++ {
				if i > begin && candidates[i] == candidates[i-1] {
					continue
				}
				tmp = append(tmp, candidates[i])
				backtrack(i+1, t-candidates[i])
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtrack(0, target)
	return result
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
