package main

import "fmt"

// https://codeforces.com/problemset/problem/1873/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Scan(&s)
		if s == "abc" || s == "acb" || s == "bac" || s == "cba" {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
