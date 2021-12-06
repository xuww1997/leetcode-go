package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	used := make([]bool, len(nums))
	tmp := make([]int, 0)
	var arrange func(begin int)
	arrange = func(begin int) {
		if begin == len(nums) {
			r := make([]int, len(tmp))
			copy(r, tmp)
			result = append(result, r)
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if !used[i] {
				tmp = append(tmp, nums[i])
				used[i] = true
				arrange(begin + 1)
				used[i] = false
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	arrange(0)
	return result
}

func main() {
	fmt.Println(permuteUnique([]int{1, 2, 1}))
}
