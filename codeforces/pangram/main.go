package main

import (
	"fmt"
	"strings"
)

// https://codeforces.com/problemset/problem/520/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)
	var s string
	_, _ = fmt.Scan(&s)

	m := make(map[rune]bool)

	for _, w := range strings.ToLower(s) {
		m[w] = true
	}

	if len(m) == 26 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
