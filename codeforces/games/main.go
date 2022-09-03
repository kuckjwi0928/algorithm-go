package main

import "fmt"

// https://codeforces.com/problemset/problem/268/A
func main() {
	var n int

	_, _ = fmt.Scan(&n)

	count := 0

	h, a := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&h[i], &a[i])
		for j := 0; j < i; j++ {
			if h[i] == a[j] {
				count++
			}
			if a[i] == h[j] {
				count++
			}
		}
	}

	fmt.Println(count)
}
