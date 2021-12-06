package main

import "fmt"

//func permute(nums []int) [][]int {
//	result := make([][]int, 0)
//	result = arrange(result, len(nums), 0, nums)
//	return result
//}
//
//func arrange(result [][]int, n int, first int, nums []int) [][]int {
//	if first == n {
//		copyNums := make([]int, n)
//		copy(copyNums, nums)
//		result = append(result, copyNums)
//	}
//	for i := first; i < n; i++ {
//		nums[i], nums[first] = nums[first], nums[i]
//		result = arrange(result, n, first+1, nums)
//		nums[i], nums[first] = nums[first], nums[i]
//	}
//	return result
//}

func permute(nums []int) [][]int {
	used := make([]bool, len(nums))
	result := make([][]int, 0)
	tmp := make([]int, 0)
	var arrange func(begin int)
	arrange = func(begin int) {
		if len(tmp) == len(nums) {
			r := make([]int, len(tmp))
			copy(r, tmp)
			result = append(result, r)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				tmp = append(tmp, nums[i])
				used[i] = true
				arrange(i + 1)
				used[i] = false
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	arrange(0)
	return result
}

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
