package main

import "fmt"

// https://codeforces.com/problemset/problem/1520/B
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for t > 0 {
		var n int
		_, _ = fmt.Scan(&n)

		count := 0

		for i := 1; i <= 9; i++ {
			for j := i; j <= n; j = j*10 + i {
				count++
			}
		}

		fmt.Println(count)

		t--
	}
}
