package main

import "fmt"

// https://codeforces.com/problemset/problem/996/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	dollars := []int{100, 20, 10, 5, 1}
	value := 0

	for _, dollar := range dollars {
		value = value + (n / dollar)
		n %= dollar
	}

	fmt.Println(value)
}
