package main

func maxArea(height []int) int {
	result := 0
	left := 0
	right := len(height) - 1
	for left < right {
		sum := 0
		if height[left] >= height[right] {
			sum = height[right] * (right - left)
			right--
		} else {
			sum = height[left] * (right - left)
			left++
		}
		if result < sum {
			result = sum
		}
	}
	return result
}
