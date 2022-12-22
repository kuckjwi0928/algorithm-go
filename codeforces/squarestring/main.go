package main

import "fmt"

// https://codeforces.com/problemset/problem/1619/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for t > 0 {
		var s string
		_, _ = fmt.Scan(&s)
		middle := len(s) / 2
		if s[:middle] == s[middle:] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		t--
	}
}
