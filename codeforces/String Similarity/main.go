package main

import "fmt"

// https://codeforces.com/problemset/problem/1400/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		var s string
		_, _ = fmt.Scan(&n, &s)
		for j := 0; j < 2*n-1; j += 2 {
			fmt.Print(string(s[j]))
		}
		fmt.Println()
	}
}
