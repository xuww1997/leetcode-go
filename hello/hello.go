package main

import (
	"fmt"
	"math"
)

func main() {
	// ascll 码
	//i := 54
	// 转换成 rune 字符类型，但是打印出来发现依然是数字样式
	//var r rune = rune(i)
	// 真正可以输出字符
	//var str string = string(rune(i))
	fmt.Println('1')
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
