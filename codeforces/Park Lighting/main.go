package main

import "fmt"

// https://codeforces.com/problemset/problem/1358/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, m int
		_, _ = fmt.Scan(&n, &m)
		total := n * m
		base := total - m
		remain := total - base
		if remain%2 != 0 {
			fmt.Println((base / 2) + ((remain / 2) + 1))
		} else {
			fmt.Println((base / 2) + (remain / 2))
		}
	}
}
