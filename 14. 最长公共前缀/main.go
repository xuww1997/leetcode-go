package main

func longestCommonPrefix(strs []string) string {
	sameLen := 0
r:
	for i := 0; ; i++ {
		for j := 0; j < len(strs); j++ {
			if i > len(strs[j])-1 {
				break r
			}
			b := strs[0][i]
			if b != strs[j][i] {
				break r
			}
		}
		sameLen++
	}
	return strs[0][:sameLen]
}
