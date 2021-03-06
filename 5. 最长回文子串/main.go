package main

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
