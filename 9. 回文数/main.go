package main

//func isPalindrome(x int) bool {
//	s := strconv.Itoa(x)
//	left := 0
//	right := len(s) - 1
//	for left <= right {
//		if s[left] != s[right] {
//			return false
//		}
//		left++
//		right--
//	}
//	return true
//}

//func isPalindrome(x int) bool {
//	if x < 0 {
//		return false
//	}
//	l := list.List{}
//	for x != 0 {
//		l.PushBack(x % 10)
//		x = x / 10
//	}
//	for l.Len() > 1 {
//		front := l.Front()
//		back := l.Back()
//		if  front.Value != back.Value {
//			return false
//		}
//		l.Remove(front)
//		l.Remove(back)
//	}
//	return true
//}

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	y := 0
	for x > y {
		y = y*10 + x%10
		x = x / 10
	}
	return x == y || x == y/10
}
