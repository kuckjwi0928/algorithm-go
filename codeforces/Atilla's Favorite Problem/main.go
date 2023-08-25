package main

import "fmt"

// https://codeforces.com/problemset/problem/1760/B
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var (
			n int
			s string
		)
		_, _ = fmt.Scan(&n, &s)
		maxKnow := 0
		for _, w := range s {
			if maxKnow < int(w-'a') {
				maxKnow = int(w - 'a')
			}
		}
		fmt.Println(maxKnow + 1)
	}
}
