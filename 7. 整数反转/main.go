package main

import "math"

//func reverse(x int) int {
//	var result int
//	for x != 0 {
//		r := x % 10
//		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && r > 7) {
//			return 0
//		}
//		if result < math.MinInt32/10 || (result == math.MinInt32/10 && r < -8) {
//			return 0
//		}
//		result = 10*result + r
//		x = x / 10
//	}
//	return result
//}
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
