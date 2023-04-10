package main

import "fmt"

// https://codeforces.com/problemset/problem/1807/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b, c int
		_, _ = fmt.Scan(&a, &b, &c)
		if a+b == c {
			fmt.Println("+")
		} else {
			fmt.Println("-")
		}
	}
}
