package main

import "fmt"

// https://codeforces.com/problemset/problem/116/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	people := 0
	m := 0

	for n != 0 {
		var a, b int
		_, _ = fmt.Scan(&a, &b)
		people = people + b - a
		m = max(m, people)
		n--
	}

	fmt.Println(m)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
