package main

import "fmt"

// https://codeforces.com/problemset/problem/1729/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b, c int
		_, _ = fmt.Scan(&a, &b, &c)

		c1 := a
		c2 := abs(b-c) + c

		if c1 == c2 {
			fmt.Println("3")
		} else if c1 < c2 {
			fmt.Println("1")
		} else {
			fmt.Println("2")
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
