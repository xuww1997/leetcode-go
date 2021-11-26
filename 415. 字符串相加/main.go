package main

import "strconv"

func addStrings(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	if len1 < len2 {
		return addStrings(num2, num1)
	}
	r := 0
	for i := 0; i < len1; i++ {
		var x int = int(num1[len1-i-1] - '0')
		var y int
		if i < len2 {
			y = int(num2[len2-i-1] - '0')
		} else {
			y = 0
		}
		sum := x + y + r
		if sum == x {
			break
		}
		r = sum / 10
		num1 = num1[:len1-i-1] + strconv.Itoa(sum%10) + num1[len1-i:]
	}
	if r != 0 {
		num1 = strconv.Itoa(r) + num1
	}
	return num1
}
