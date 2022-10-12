package main

import "fmt"

// https://codeforces.com/problemset/problem/432/A
func main() {
	var n, k int
	_, _ = fmt.Scan(&n, &k)

	count := 0

	for i := 0; i < n; i++ {
		var j int
		_, _ = fmt.Scan(&j)
		if j+k <= 5 {
			count++
		}
	}

	fmt.Println(count / 3)
}
