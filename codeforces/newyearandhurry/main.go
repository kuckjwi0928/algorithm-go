package main

import "fmt"

// https://codeforces.com/problemset/problem/750/A
func main() {
	var n, k int
	_, _ = fmt.Scan(&n, &k)

	deadline := 60*4 - k
	sum := 0
	count := 0

	for i := 1; i <= n; i++ {
		sum += 5 * i
		if sum > deadline {
			break
		}
		count++
	}

	fmt.Println(count)
}
