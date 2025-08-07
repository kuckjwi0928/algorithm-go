package main

import "fmt"

// https://codeforces.com/problemset/problem/1999/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)
		first, second := toDigits(n)
		fmt.Println(first + second)
	}
}

func toDigits(n int) (int, int) {
	first := n % 10
	second := n / 10
	return first, second
}
