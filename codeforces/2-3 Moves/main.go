package main

import (
	"fmt"
)

// https://codeforces.com/problemset/problem/1716/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)
		if n == 1 {
			fmt.Println(2)
		} else if n%3 != 0 {
			fmt.Println(n/3 + 1)
		} else {
			fmt.Println(n / 3)
		}
	}
}
