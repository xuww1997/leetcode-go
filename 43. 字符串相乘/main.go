package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(multiply("123", "456"))
}
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	result := ""
	len1 := len(num1)
	len2 := len(num2)
	arr := make([]int, len1+len2)
	for i := len1 - 1; i >= 0; i-- {
		var x int = int(num1[i] - '0')
		for j := len2 - 1; j >= 0; j-- {
			var y int = int(num2[j] - '0')
			arr[i+j+1] += x * y
		}
	}
	for i := len1 + len2 - 1; i > 0; i-- {
		arr[i-1] += arr[i] / 10
		arr[i] %= 10
	}
	i := 0
	if arr[0] == 0 {
		i = 1
	}
	for ; i < len1+len2; i++ {
		result = result + strconv.Itoa(arr[i])
	}
	return result
}

//func multiply(num1 string, num2 string) string {
//	if num1 == "0" || num2 == "0" {
//		return "0"
//	}
//	result := "0"
//	len1 := len(num1)
//	len2 := len(num2)
//	for i := 0; i < len1; i++ {
//		var x int = int(num1[len1-i-1] - '0')
//
//		for j := 0; j < len2; j++ {
//			var y int = int(num2[len2-j-1] - '0')
//			z := strconv.Itoa(x * y)
//			for k := 0; k < i+j; k++ {
//				z = z + "0"
//			}
//			result = addStrings(result, z)
//		}
//	}
//	return result
//}
//func addStrings(num1 string, num2 string) string {
//	len1 := len(num1)
//	len2 := len(num2)
//	if len1 < len2 {
//		return addStrings(num2, num1)
//	}
//	r := 0
//	for i := 0; i < len1; i++ {
//		var x int = int(num1[len1-i-1] - '0')
//		var y int
//		if i < len2 {
//			y = int(num2[len2-i-1] - '0')
//		} else {
//			y = 0
//		}
//		sum := x + y + r
//		if sum == x && i >= len2 {
//			break
//		}
//		r = sum / 10
//		num1 = num1[:len1-i-1] + strconv.Itoa(sum%10) + num1[len1-i:]
//	}
//	if r != 0 {
//		num1 = strconv.Itoa(r) + num1
//	}
//	return num1
//}
