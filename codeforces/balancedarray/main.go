package main

import "fmt"

// https://codeforces.com/problemset/problem/151/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)
		if n%4 != 0 {
			fmt.Println("NO")
			continue
		}
		fmt.Println("YES")
		for j := 1; j <= n/2; j++ {
			fmt.Println(j * 2)
		}
		for j := 1; j < n/2; j++ {
			fmt.Println(j*2 - 1)
		}
		fmt.Println(n + n/2 - 1)
	}
}
