package main

import "fmt"

// https://codeforces.com/problemset/problem/1791/A
func main() {
	m := map[string]bool{
		"c": true,
		"o": true,
		"d": true,
		"e": true,
		"f": true,
		"s": true,
		"r": true,
	}

	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var c string
		_, _ = fmt.Scan(&c)
		if m[c] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
