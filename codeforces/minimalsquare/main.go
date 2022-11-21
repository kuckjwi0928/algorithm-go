package main

import "fmt"

// https://codeforces.com/problemset/problem/1360/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for t > 0 {
		var a, b int
		_, _ = fmt.Scan(&a, &b)
		if a > b {
			a, b = swap(a, b)
		}
		fmt.Println(max(a*2, b) * max(a*2, b))
		t--
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func swap(a, b int) (int, int) {
	return b, a
}
