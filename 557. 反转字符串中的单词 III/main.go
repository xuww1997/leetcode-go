package main

import (
	"fmt"
)

//func reverseWords(s string) string {
//	if s == "" {
//		return s
//	}
//	result := ""
//	arr := strings.Split(s, " ")
//	for _, str := range arr {
//		reverseStr := make([]byte,len(str))
//		for i := 0;i<len(str);i++{
//			reverseStr[i] = str[len(str)-1-i]
//		}
//		result = result + string(reverseStr) + " "
//	}
//	return result[:len(result)-1]
//}

func reverseWords(s string) string {
	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); {
		j, start := i, i
		for i < len(s) && s[i] != ' ' {
			i++
		}
		for j < i {
			res = append(res, s[i-j-1+start])
			j++
		}
		if i < len(s) && s[i] == ' ' {
			res = append(res, ' ')
			i++
		}
	}
	return string(res)
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
