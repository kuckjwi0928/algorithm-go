package main

import "fmt"

// https://codeforces.com/problemset/problem/71/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Scan(&s)
		if len(s) > 10 {
			fmt.Print(string(s[0]))
			fmt.Print(len(s[1 : len(s)-1]))
			fmt.Print(string(s[len(s)-1]))
			fmt.Println()
		} else {
			fmt.Println(s)
		}
	}
}
