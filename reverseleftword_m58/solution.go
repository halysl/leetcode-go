package reverseleftword_m58

import "fmt"

func reverseLeftWords(s string, n int) string {
	if n >= len(s) {
		return s
	}
	return s[n:]+s[:n]
}

func Run() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
}