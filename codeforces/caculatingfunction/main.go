package main

import (
	"fmt"
)

// https://codeforces.com/problemset/problem/486/A
func main() {
	var n int64
	_, _ = fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println(n / 2)
	} else {
		fmt.Println(((n + 1) / 2) * -1)
	}
}
