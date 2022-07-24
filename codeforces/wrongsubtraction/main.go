package main

import "fmt"

// https://codeforces.com/problemset/problem/977/A
func main() {
	var n, k int
	_, _ = fmt.Scan(&n, &k)

	for k != 0 {
		if n%10 == 0 {
			n /= 10
		} else {
			n--
		}
		k--
	}

	fmt.Println(n)
}
