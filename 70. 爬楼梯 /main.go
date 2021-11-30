package main

import "fmt"

//func climbStairs(n int) int {
//	if n < 2 {
//		return n
//	}
//	arr := make([]int, n)
//	arr[0] = 1
//	arr[1] = 2
//	for i := 2; i < n; i++ {
//		arr[i] = arr[i-1] + arr[i-2]
//	}
//	return arr[n-1]
//}

func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	s1, s2 := 1, 2
	for i := 2; i < n; i++ {
		s1, s2 = s2, s1+s2
	}
	return s2
}

func main() {
	fmt.Println(climbStairs(4))
}
