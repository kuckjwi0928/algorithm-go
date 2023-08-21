package main

import "fmt"

// https://codeforces.com/problemset/problem/1850/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var a, b, c int
		_, _ = fmt.Scan(&a, &b, &c)
		if a+b >= 10 || b+c >= 10 || a+c >= 10 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
