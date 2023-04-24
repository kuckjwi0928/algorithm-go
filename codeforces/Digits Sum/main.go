package main

import "fmt"

// https://codeforces.com/problemset/problem/1553/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)
		fmt.Println((n + 1) / 10)
	}
}
