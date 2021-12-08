package main

//func majorityElement(nums []int) int {
//	x := len(nums)/2
//	m := map[int]int{}
//	for _, v := range nums {
//		n := m[v] + 1
//		if n > x{
//			return v
//		}
//		m[v] = n
//	}
//	return 0
//}

func majorityElement(nums []int) int {
	count := 0
	num := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == num {
			count++
		} else {
			count--
		}
		if count < 0 {
			num = nums[i]
			count = 0
		}
	}
	return num
}
