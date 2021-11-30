package main

import "fmt"

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	result = arrange(result, len(nums), 0, nums)
	return result
}

func arrange(result [][]int, n int, first int, nums []int) [][]int {
	if first == n {
		copyNums := make([]int, n)
		copy(copyNums, nums)
		result = append(result, copyNums)
	}
	for i := first; i < n; i++ {
		nums[i], nums[first] = nums[first], nums[i]
		result = arrange(result, n, first+1, nums)
		nums[i], nums[first] = nums[first], nums[i]
	}
	return result
}

//
//func permute(nums []int) [][]int {
//	res = [][]int{}
//	run(nums, []int{})
//	return res
//}
//
//var res [][]int
//
//func run(nums []int, cur []int) {
//	if len(cur) == len(nums) {
//		res = append(res, cur)
//		return
//	}
//	for i := 0; i < len(nums); i++ {
//		exist := false
//		for j := 0; j < len(cur); j++ {
//			if nums[i] == cur[j] {
//				exist = true
//				break
//			}
//		}
//		if exist {
//			continue
//		}
//		run(nums, append(cur, nums[i]))
//	}
//}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
	//result := make([][]int, 0)
	//result = addResult(result)
	//fmt.Println(result)
}

func addResult(result [][]int) [][]int {
	result = append(result, []int{1, 2, 3})
	return result
}
