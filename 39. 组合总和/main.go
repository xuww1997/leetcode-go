package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	tmp := make([]int, 0)
	var backtrack func(begin int, t int)
	backtrack = func(begin int, t int) {
		if t == 0 {
			r := make([]int, len(tmp))
			copy(r, tmp)
			result = append(result, r)
		} else if t < 0 {
			return
		} else {
			for i := begin; i < len(candidates); i++ {
				tmp = append(tmp, candidates[i])
				backtrack(i, t-candidates[i])
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtrack(0, target)
	return result
}

//func combinationSum(candidates []int, target int) [][]int {
//	result := make([][]int, 0)
//	result = backtrack(candidates, target, result, 0, []int{})
//	return result
//}

//func backtrack(candidates []int, target int, result [][]int, begin int, tmp []int) [][]int {
//	if target < 0 {
//		return result
//	} else if target == 0 {
//		r := make([]int, len(tmp))
//		copy(r, tmp)
//		return append(result, r)
//	}
//	for i := begin; i < len(candidates); i++ {
//		tmp = append(tmp, candidates[i])
//		result = backtrack(candidates, target-candidates[i], result, i, tmp)
//		tmp = tmp[:len(tmp)-1]
//	}
//	return result
//}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
