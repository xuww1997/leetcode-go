package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	//arr := []int{1,3,2}
	//sort.Ints(arr)
	fmt.Println(threeSumClosest([]int{1, 1, -1, -1, 3}, -1))
	//fmt.Println()
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := math.MaxInt32
	nearestNum := math.MaxInt32
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := len(nums) - 1
		for second := first + 1; second < third; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for second < third && nums[first]+nums[second]+nums[third] > target {
				abs := int(math.Abs(float64(nums[first] + nums[second] + nums[third] - target)))
				if nearestNum > abs {
					nearestNum = abs
					result = nums[first] + nums[second] + nums[third]
				}
				third--
			}
			if second != third && nums[first]+nums[second]+nums[third] < target {
				abs := int(math.Abs(float64(nums[first] + nums[second] + nums[third] - target)))
				if nearestNum > abs {
					nearestNum = abs
					result = nums[first] + nums[second] + nums[third]
				}
				continue
			}
			if second != third && nums[first]+nums[second]+nums[third] == target {
				return nums[first] + nums[second] + nums[third]
			}
		}
	}
	return result
}

func compare(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left]+nums[right]+nums[i] == 0 {
				result = append(result, []int{nums[left], nums[right], nums[i]})
				for left < len(nums)-2 && nums[left] == nums[left+1] {
					left++
				}
				for right > 1 && nums[right] == nums[right-1] {
					right--
				}
				right--
				left++
			} else if nums[left]+nums[right]+nums[i] < 0 {
				left++
			} else if nums[left]+nums[right]+nums[i] > 0 {
				right--
			}
		}
	}
	return result
}

func longestCommonPrefix(strs []string) string {
	sameLen := 0
r:
	for i := 0; ; i++ {
		for j := 0; j < len(strs); j++ {
			if i > len(strs[j])-1 {
				break r
			}
			b := strs[0][0]
			if b != strs[j][0] {
				break r
			}
			sameLen++
		}
	}
	return strs[0][:sameLen]
}
func isPalindrome(x int) bool {
	l := list.List{}
	for x != 0 {
		l.PushBack(x % 10)
		x = x / 10
	}
	for l.Len() > 1 {
		front := l.Front()
		back := l.Back()
		if front.Value != back.Value {
			return false
		}
		l.Remove(front)
		l.Remove(back)
	}
	return true

}
func myAtoi(s string) int {
	result := 0
	readEmpty := true
	readPositive := true
	isPositive := 1
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if !readEmpty {
				break
			}
		} else if s[i] == '-' || s[i] == '+' {
			if !readPositive {
				break
			}
			if s[i] == '-' {
				isPositive = -1
			}
			readEmpty = false
			readPositive = false
		} else if s[i] >= '0' && s[i] <= '9' {
			n, _ := strconv.Atoi(string(rune(s[i])))
			result = result*10 + n
			readEmpty = false
			readPositive = false
			if result*isPositive > math.MaxInt32 {
				return math.MaxInt32
			} else if result*isPositive < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			break
		}
	}

	return result * isPositive
}

func reverse(x int) int {
	var result int
	for x != 0 {
		result = 10*result + x%10
		x = x / 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	} else {
		return result
	}
}
func longestPalindrome(s string) string {
	var maxLen int
	var result string
	for i := 0; i < len(s); i++ {
		len1, left1 := getMaxLength(i, i, s)
		len2, left2 := getMaxLength(i, i+1, s)
		var len, left int
		if len1 >= len2 {
			len = len1
			left = left1
		} else {
			len = len2
			left = left2
		}
		if maxLen < len {
			maxLen = len
			result = s[left : left+len]
		}
	}
	return result
}

func getMaxLength(left int, right int, s string) (int, int) {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return right - left - 1, left + 1

}
