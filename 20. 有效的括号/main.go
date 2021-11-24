package main

func isValid(s string) bool {
	stack := make([]byte, 0)
	last := -1
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	add := func(c byte) {
		stack = append(stack, c)
		last++
	}
	removeLast := func() {
		stack = stack[:last]
		last--
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			add(s[i])
		} else {
			if last < 0 {
				return false
			}
			if pairs[s[i]] != stack[last] {
				return false
			}
			removeLast()
		}
	}
	return last == -1
}
