package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 || n > 0 {
		if m <= 0 {
			nums1[n-1] = nums2[n-1]
			n--
			continue
		} else if n <= 0 {
			break
		}
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	for i := 0; i < len(nums1); i++ {
		fmt.Println(nums1[i])
	}
}

func main() {
	merge([]int{0}, 0, []int{1}, 1)
}
