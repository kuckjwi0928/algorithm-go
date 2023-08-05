package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://codeforces.com/problemset/problem/1800/A
func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		var s string
		_, _ = fmt.Fscan(reader, &n)
		_, _ = fmt.Fscan(reader, &s)
		s = strings.ToLower(s)
		words := make([]byte, 0)
		for j := 0; j < len(s)-1; j++ {
			if s[j] != s[j+1] {
				words = append(words, s[j])
			}
		}
		words = append(words, s[len(s)-1])
		if "meow" == string(words) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
