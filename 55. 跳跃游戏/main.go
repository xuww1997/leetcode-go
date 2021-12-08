package main

func canJump(nums []int) bool {
	k := 0
	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}
		if k < i+nums[i] {
			k = i + nums[i]
		}
	}
	return true
}
