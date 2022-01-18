package main

//func productExceptSelf(nums []int) []int {
//	res, L, R := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
//	L[0] = 1
//	R[len(nums)-1] = 1
//	for i := 1; i < len(nums); i++ {
//		L[i] = nums[i-1] * L[i-1]
//	}
//	for i := len(nums) - 2; i >= 0; i-- {
//		R[i] = nums[i+1] * R[i+1]
//	}
//	for i := 0; i < len(nums); i++ {
//		res[i] = L[i] * R[i]
//	}
//	return res
//}

//func productExceptSelf(nums []int) []int {
//	res := make([]int, len(nums))
//	res[0] = 1
//	for i := 1; i < len(nums); i++ {
//		res[i] = res[i-1] * nums[i-1]
//	}
//	R := 1
//	for i := len(nums) - 1; i >= 0; i-- {
//		res[i] = res[i] * R
//		R *= nums[i]
//	}
//	return res
//}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	R := 1
	L := 1
	for i := 0; i < len(nums); i++ {
		res[i] = L
		L *= nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = R * nums[i]
		R *= nums[i]
	}
	return res
}
