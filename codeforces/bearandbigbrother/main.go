package main

import "fmt"

// https://codeforces.com/problemset/problem/791/A
func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)
	year := 0
	for a <= b {
		a = a * 3
		b = b * 2
		year++
	}
	fmt.Println(year)
}
