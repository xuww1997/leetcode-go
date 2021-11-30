package main

import "fmt"

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	left, right, up, down := 0, n-1, 0, n-1
	num := 1
	for num <= n*n {
		for y := left; y <= right; y++ {
			result[up][y] = num
			num++
		}
		up++
		for x := up; x <= down; x++ {
			result[x][right] = num
			num++
		}
		right--
		for y := right; y >= left; y-- {
			result[down][y] = num
			num++
		}
		down--
		for x := down; x >= up; x-- {
			result[x][left] = num
			num++
		}
		left++
	}
	return result
}

func main() {
	fmt.Println(generateMatrix(3))
}
