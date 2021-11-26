package main

//func removeDuplicates(nums []int) int {
//	newLen := 0
//	numMap := map[int]int{}
//	for i := 0; i < len(nums); i++ {
//		if numMap[nums[i]] == 0 {
//			numMap[nums[i]] = 1
//			nums[newLen] = nums[i]
//			newLen++
//		}
//	}
//	return newLen
//}

func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
