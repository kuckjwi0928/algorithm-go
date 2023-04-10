package main

import (
	"fmt"
)

// https://codeforces.com/problemset/problem/1814/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, k int
		_, _ = fmt.Scan(&n, &k)
		if k%2 == 0 {
			if n%2 == 0 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		} else {
			fmt.Println("YES")
		}
	}
}
