package main

import "fmt"

// https://codeforces.com/problemset/problem/467/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	count := 0

	for i := 0; i < n; i++ {
		var p, q int
		_, _ = fmt.Scan(&p, &q)
		if q-p >= 2 {
			count++
		}
	}

	fmt.Println(count)
}
