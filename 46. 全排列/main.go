package main

import "fmt"

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		remainArr := append(nums[:i], nums[i+1:]...)
		result = append(result, arrange([]int{nums[i]}, remainArr)...)
	}
	return result
}

func arrange(nums1 []int, nums2 []int) [][]int {
	if len(nums2) == 1 {
		return [][]int{[]int{nums1[0], nums2[0]}, []int{nums2[0], nums1[0]}}
	}
	arr := arrange(nums2[:1], nums2[1:])
	for _, a := range arr {
		a = append(nums1, a...)
	}
	return arr

}

func main() {
	var r float64 = 150
	k := 0.2
	for i := 0; i < 8; i++ {
		r = r + (1+k)*40.8
		k += 0.2
	}
	r = r + 321*(1+k)
	fmt.Println(r)
}
