package main

import "fmt"

// https://codeforces.com/problemset/problem/492/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)
	height := 1
	for n > 0 {
		sum := 0
		for i := 1; i <= height; i++ {
			sum += i
		}
		n -= sum
		height++
	}
	if n == 0 {
		fmt.Println(height - 1)
	} else {
		fmt.Println(height - 2)
	}
}
