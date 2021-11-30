package main

import "fmt"

//func uniquePaths(m int, n int) int {
//	return getPath(m, n)
//}
//
//func getPath(m int, n int) int {
//	if m == 1 {
//		return 1
//	}else if n == 1 {
//		return 1
//	}else {
//		return getPath(m-1, n) + getPath(m, n-1)
//	}
//}

//func uniquePaths(m int, n int) int {
//	arr := make([][]int, m)
//	for i := range arr {
//		arr[i] = make([]int, n)
//		arr[i][0] = 1
//	}
//	for j := 0; j < n; j++ {
//		arr[0][j] = 1
//	}
//	for i := 1; i < m; i++ {
//		for j := 1; j < n; j++ {
//			arr[i][j] = arr[i-1][j] + arr[i][j-1]
//		}
//	}
//	return arr[m-1][n-1]
//}

func uniquePaths(m int, n int) int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			arr[j] = arr[j] + arr[j-1]
		}
	}
	return arr[n-1]
}
func main() {
	fmt.Println(uniquePaths(51, 9))
}
