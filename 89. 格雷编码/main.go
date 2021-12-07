package main

import "fmt"

func grayCode(n int) []int {
	result := []int{0}
	head := 1
	for ; n > 0; n-- {
		for i := len(result) - 1; i >= 0; i-- {
			result = append(result, head+result[i])
		}
		head <<= 1
	}
	return result
}

func main() {
	fmt.Println(grayCode(3))
}
