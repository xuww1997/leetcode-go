package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	tmp := make([]int, 0)
	var backtrack func(total int, begin int)
	backtrack = func(total int, begin int) {
		if len(tmp) == total {
			r := make([]int, len(tmp))
			copy(r, tmp)
			result = append(result, r)
		} else {
			for i := begin; i < len(nums); i++ {
				if i > begin && nums[i] == nums[i-1] {
					continue
				}
				tmp = append(tmp, nums[i])
				backtrack(total, i+1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	for i := 0; i <= len(nums); i++ {
		backtrack(i, 0)
		tmp = make([]int, 0)
	}
	return result
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
