package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0)
	left, right, up, down := 0, len(matrix[0])-1, 0, len(matrix)-1
	for left < right && up < down {
		for y := left; y <= right; y++ {
			result = append(result, matrix[up][y])
		}
		up++
		for x := up; x <= down; x++ {
			result = append(result, matrix[x][right])
		}
		right--
		for y := right; y >= left; y-- {
			result = append(result, matrix[down][y])
		}
		down--
		for x := down; x >= up; x-- {
			result = append(result, matrix[x][left])
		}
		left++
	}
	if up == down {
		for y := left; y <= right; y++ {
			result = append(result, matrix[up][y])
		}
	} else if left == right {
		for x := up; x <= down; x++ {
			result = append(result, matrix[x][right])
		}
	}
	return result
}

//func circle(left int, right int, up int, down int, result []int, matrix [][]int) []int {
//	for y := left; y <= right; y++ {
//		result = append(result, matrix[up][y])
//	}
//	for x := up + 1; x <= down; x++ {
//		result = append(result, matrix[x][right])
//	}
//	for y := right - 1; y >= left; y-- {
//		result = append(result, matrix[down][y])
//	}
//	for x := down - 1; x >= up; x-- {
//		result = append(result, matrix[x][left])
//	}
//}

func main() {
	//arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	arr := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	//fmt.Println(len(arr))
	//fmt.Println(arr[0][0])
	//fmt.Println(arr[0][1])
	//fmt.Println(arr[0][2])
	fmt.Println(spiralOrder(arr))
}
