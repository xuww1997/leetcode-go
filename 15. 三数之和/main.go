package main

import "sort"

//func threeSum(nums []int) [][]int {
//	result := make([][]int,0)
//	if len(nums) < 3 {
//		return result
//	}
//	sort.Ints(nums)
//	for i := 0; i < len(nums); i++ {
//		if nums[i] > 0 {
//			break
//		}
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		left := i + 1
//		right := len(nums) - 1
//		for left < right {
//			if nums[left]+nums[right]+nums[i] == 0 {
//				result = append(result, []int{nums[left], nums[right], nums[i]})
//				for left < len(nums)-2 && nums[left] == nums[left+1] {
//					left++
//				}
//				for right > 1 && nums[right] == nums[right-1] {
//					right--
//				}
//				right--
//				left++
//			} else if nums[left]+nums[right]+nums[i] < 0 {
//				left++
//			} else if nums[left]+nums[right]+nums[i] > 0 {
//				right--
//			}
//		}
//	}
//	return result
//}

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for first := 0; first < len(nums); first++ {
		if nums[first] > 0 {
			break
		}
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := len(nums) - 1
		for second := first + 1; second < third; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[first]+nums[second]+nums[third] > 0 {
				third--
			}
			if second != third && nums[first]+nums[second]+nums[third] == 0 {
				result = append(result, []int{nums[first], nums[second], nums[third]})
			}
		}

	}
	return result
}
