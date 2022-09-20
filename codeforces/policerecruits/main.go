package main

import "fmt"

// https://codeforces.com/problemset/problem/427/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	available := 0
	crime := 0

	for i := 0; i < n; i++ {
		var j int
		_, _ = fmt.Scan(&j)

		if j == -1 && available == 0 {
			crime++
		} else if j == -1 && available > 0 {
			available--
		} else {
			available += j
		}
	}

	fmt.Println(crime)
}
