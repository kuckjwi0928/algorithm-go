package main

import (
	"fmt"
)

// https://codeforces.com/problemset/problem/978/B
func main() {
	var n int
	var s string

	_, _ = fmt.Scan(&n)
	_, _ = fmt.Scan(&s)

	count := 0

	for i := 0; i < n; i++ {
		if i+2 < n {
			if (s[i] == 'x') && (s[i+1] == 'x') && (s[i+2] == 'x') {
				count++
			}
		}
	}

	fmt.Println(count)
}
