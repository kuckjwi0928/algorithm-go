package main

import "fmt"

// https://codeforces.com/problemset/problem/1328/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var a, b int
		_, _ = fmt.Scan(&a, &b)

		if a%b == 0 {
			fmt.Println("0")
		} else {
			fmt.Println(b - (a % b))
		}
	}
}
