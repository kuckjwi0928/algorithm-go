package main

import "fmt"

// https://codeforces.com/problemset/problem/1367/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Scan(&s)
		for j := 0; j < len(s); j += 2 {
			fmt.Print(string(s[j]))
		}
		fmt.Print(string(s[len(s)-1]))
		fmt.Println()
	}
}
