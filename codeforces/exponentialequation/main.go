package main

import (
	"fmt"
)

// https://codeforces.com/problemset/problem/1787/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int64
		_, _ = fmt.Scan(&n)
		if n%2 == 0 {
			fmt.Printf("%d 1\n", n/2)
		} else {
			fmt.Println(-1)
		}
	}
}
