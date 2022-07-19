package main

import "fmt"

// https://codeforces.com/problemset/problem/617/A
func main() {
	var x int
	_, _ = fmt.Scan(&x)
	count := x / 5
	if x%5 > 0 {
		count++
	}
	fmt.Println(count)
}
