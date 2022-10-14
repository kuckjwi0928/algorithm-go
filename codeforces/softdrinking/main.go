package main

import "fmt"

// https://codeforces.com/problemset/problem/151/A
func main() {
	var n, k, l, c, d, p, nl, np int
	_, _ = fmt.Scan(&n, &k, &l, &c, &d, &p, &nl, &np)
	fmt.Println(min(min(k*l/nl, c*d), p/np) / n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
