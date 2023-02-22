package main

import "fmt"

// https://codeforces.com/problemset/problem/1353/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, m int64
		_, _ = fmt.Scan(&n, &m)
		if n == 1 {
			fmt.Println(0)
		} else if n == 2 {
			fmt.Println(m)
		} else {
			fmt.Println(m * 2)
		}
	}
}
