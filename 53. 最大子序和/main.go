package main

import "fmt"

//func maxSubArray(nums []int) int {
//	sum := nums[0]
//	max := nums[0]
//	for i := 1; i < len(nums); i++ {
//		if nums[i] > sum + nums[i] {
//			sum = nums[i]
//		}else {
//			sum += nums[i]
//		}
//		if max < sum {
//			max = sum
//		}
//	}
//	return max
//}
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i]+nums[i-1] {
			nums[i] = nums[i] + nums[i-1]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
