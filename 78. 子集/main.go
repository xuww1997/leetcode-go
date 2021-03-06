package main

import (
	"fmt"
)

//func subsets(nums []int) [][]int {
//	result := make([][]int,0)
//	for mask := 0; mask < 1<<len(nums); mask++ {
//		r := make([]int, 0)
//		for i := 0; i < len(nums); i++ {
//			if mask>>i&1 == 1 {
//				r = append(r, nums[i])
//			}
//		}
//		result = append(result, r)
//	}
//	return result
//}

func subsets(nums []int) [][]int {
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
	fmt.Println(subsets([]int{1, 2, 3}))
}
