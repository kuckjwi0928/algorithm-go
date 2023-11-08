package main

import "fmt"

// https://leetcode.com/problems/reverse-words-in-a-string-ii/description/
func main() {
	s := []byte("the sky is blue")
	reverseWords(s)
	fmt.Println(string(s))
}

func reverseWords(s []byte) {
	reverse(s, 0, len(s)-1)
	for i, j := 0, 0; j <= len(s); j++ {
		if j == len(s) || s[j] == ' ' {
			reverse(s, i, j-1)
			i = j + 1
		}
	}
}

func reverse(s []byte, i, j int) {
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
