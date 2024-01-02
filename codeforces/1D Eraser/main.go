package main

import "fmt"

// https://codeforces.com/problemset/problem/1873/D
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, k, count int
		_, _ = fmt.Scan(&n, &k)
		var s string
		_, _ = fmt.Scan(&s)
		for j := 0; j < n; j++ {
			if s[j] == 'B' {
				count++
				j += k - 1
			}
		}
		fmt.Println(count)
	}
}
