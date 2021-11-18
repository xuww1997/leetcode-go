package main

import (
	"math"
)

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
			//n, _ := strconv.Atoi(string(rune(s[i])))
			result = result*10 + int(s[i]-'0')
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
