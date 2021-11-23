package main

import (
	"math"
	"sort"
)

//func threeSumClosest(nums []int, target int) int {
//	sort.Ints(nums)
//	result := math.MaxInt32
//	nearestNum := math.MaxInt32
//	for first := 0; first < len(nums); first++ {
//		if first > 0 && nums[first] == nums[first-1] {
//			continue
//		}
//		third := len(nums) - 1
//		for second := first + 1; second < third; second++ {
//			if second > first+1 && nums[second] == nums[second-1] {
//				continue
//			}
//
//			for second < third && nums[first] + nums[second] + nums[third] > target {
//				abs := int(math.Abs(float64(nums[first] + nums[second] + nums[third] - target)))
//				if nearestNum > abs {
//					nearestNum = abs
//					result = nums[first] + nums[second] + nums[third]
//				}
//				third--
//			}
//			if second != third && nums[first] + nums[second] + nums[third] < target {
//				abs := int(math.Abs(float64(nums[first] + nums[second] + nums[third] - target)))
//				if nearestNum > abs {
//					nearestNum = abs
//					result = nums[first] + nums[second] + nums[third]
//				}
//				continue
//			}
//			if second != third && nums[first] + nums[second] + nums[third] == target {
//				return nums[first] + nums[second] + nums[third]
//			}
//		}
//	}
//	return result
//}
func threeSumClosest(nums []int, target int) int {
	result := math.MaxInt32
	sort.Ints(nums)
	update := func(n int) {
		if math.Abs(float64(n-target)) < math.Abs(float64(result-target)) {
			result = n
		}
	}
	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return target
			} else if sum > target {
				r--
				for r > l && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < target {
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
			update(sum)
		}
	}
	return result
}
