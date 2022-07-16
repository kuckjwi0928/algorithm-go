package main

import "fmt"

// https://codeforces.com/problemset/problem/546/A
func main() {
	var k, n, w int
	_, _ = fmt.Scan(&k, &n, &w)

	for i := 1; i <= w; i++ {
		pay := i * k
		n -= pay
	}

	if n < 0 {
		fmt.Println(-n)
	} else {
		fmt.Println(0)
	}
}
