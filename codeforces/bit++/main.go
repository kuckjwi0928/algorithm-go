package main

import (
	"fmt"
)

// https://codeforces.com/problemset/problem/282/A
func main() {
	var n int
	var s string
	_, _ = fmt.Scan(&n)
	x := 0
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&s)
		if s[1] == '+' {
			x++
		} else {
			x--
		}
	}
	fmt.Println(x)
}
