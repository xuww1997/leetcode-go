package main

import "fmt"

//func isPowerOfTwo(n int) bool {
//	if n <= 0{
//		return false
//	}
//	for n != 1 {
//		if n%2 != 0 {
//			return false
//		}
//		n >>= 1
//	}
//	return true
//}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func main() {
	fmt.Println(isPowerOfTwo(16))
}
