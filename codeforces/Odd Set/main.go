package main

import "fmt"

// https://codeforces.com/problemset/problem/1542/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)

		count := 0

		for j := 0; j < n*2; j++ {
			var a int
			_, _ = fmt.Scan(&a)
			count += a % 2
		}

		if count == n {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
