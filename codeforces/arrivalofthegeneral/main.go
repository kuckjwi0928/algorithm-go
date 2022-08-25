package main

import "fmt"

// https://codeforces.com/problemset/problem/144/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	min := 1000
	minIdx := 0
	max := 0
	maxIdx := 0

	for i := 0; i < n; i++ {
		var j int
		_, _ = fmt.Scan(&j)
		if j > max {
			max = j
			maxIdx = i
		}
		if j <= min {
			min = j
			minIdx = i
		}
	}

	if minIdx < maxIdx {
		fmt.Println(maxIdx + (n - 1 - minIdx) - 1)
	} else {
		fmt.Println(maxIdx + (n - 1 - minIdx))
	}
}
