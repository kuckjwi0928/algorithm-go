package main

import "fmt"

// https://codeforces.com/problemset/problem/707/A
func main() {
	var n, m int
	_, _ = fmt.Scan(&n, &m)
	color := false
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var c string
			_, _ = fmt.Scan(&c)
			if c == "C" || c == "M" || c == "Y" {
				color = true
				break
			}
		}
	}
	if color {
		fmt.Println("#Color")
	} else {
		fmt.Println("#Black&White")
	}
}
