package main

import "fmt"

// https://codeforces.com/problemset/problem/1368/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b, n int
		_, _ = fmt.Scan(&a, &b, &n)
		count := 1
		for n >= a+b {
			if a > b {
				b += a
			} else {
				a += b
			}
			count++
		}
		fmt.Println(count)
	}
}
