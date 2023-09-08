package main

import "fmt"

// https://codeforces.com/problemset/problem/1569/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var (
			n int
			s string
		)
		_, _ = fmt.Scan(&n, &s)

		left := -1
		right := -1

		for j := 1; j < len(s); j++ {
			if s[j] != s[j-1] {
				left = j
				right = j + 1
				break
			}
		}

		fmt.Println(left, right)
	}
}
