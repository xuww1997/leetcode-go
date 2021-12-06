package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	l := len(s)
	if l < 4 || l > 12 {
		return []string{}
	}
	//pLen := l / 4
	nums := make([]string, 0)
	var backtrack func(part int, index int)
	backtrack = func(part int, index int) {
		if part > 0 && !isValid(nums, s, part, index) {
			return
		}
		if part == 4 {
			if index == l {
				result = append(result, sliceToIp(nums))
				return
			}
		}
		for i := 1; i <= 3; i++ {
			if index+i > len(s) {
				return
			}
			nums = append(nums, s[index:index+i])
			backtrack(part+1, index+i)
			nums = nums[0 : len(nums)-1]
		}
	}
	backtrack(0, 0)
	return result
}

func isValid(nums []string, s string, part int, index int) bool {
	num := nums[part-1]
	n, _ := strconv.Atoi(num)
	if (num[0:1] == "0" && len(num) > 1) || n > 255 {
		return false
	}
	remainLen := len(s[index:])
	if remainLen > 3*(4-part) || remainLen < 1*(4-part) {
		return false
	}
	return true
}

func sliceToIp(nums []string) string {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = result + "." + nums[i]
	}
	return result
}

func main() {
	fmt.Println(restoreIpAddresses("0000"))
}
