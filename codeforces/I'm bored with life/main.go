package main

import "fmt"

// https://codeforces.com/problemset/problem/822/A
func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)
	n := min(a, b)
	res := 1
	for n > 0 {
		res *= n
		n--
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
