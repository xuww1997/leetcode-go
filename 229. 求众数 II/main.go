package main

import "fmt"

func majorityElement(nums []int) []int {
	result := make([]int, 0)
	if nums == nil || len(nums) == 0 {
		return result
	}
	num1, num2 := nums[0], nums[0]
	vote1, vote2 := 0, 0
	for _, n := range nums {
		if n == num1 {
			vote1++
		} else if n == num2 {
			vote2++
		} else {
			if vote1 == 0 {
				num1 = n
				vote1++
			} else if vote2 == 0 {
				num2 = n
				vote2++
			} else {
				vote1--
				vote2--
			}
		}
	}
	vote1, vote2 = 0, 0
	for _, n := range nums {
		if n == num1 {
			vote1++
		} else if n == num2 {
			vote2++
		}
	}
	if vote1 > len(nums)/3 {
		result = append(result, num1)
	}
	if vote2 > len(nums)/3 {
		result = append(result, num2)
	}
	return result
}

func main() {
	fmt.Println(majorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2}))
}
