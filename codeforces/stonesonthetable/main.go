package main

import "fmt"

// https://codeforces.com/problemset/problem/266/A
func main() {
	var n int
	var input string

	_, _ = fmt.Scan(&n)
	_, _ = fmt.Scan(&input)

	fetch := 0

	for i := 1; i < n; i++ {
		if input[i-1] != input[i] {
			continue
		}
		fetch++
	}

	fmt.Println(fetch)
}
