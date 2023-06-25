package main

import "fmt"

// https://codeforces.com/problemset/problem/1675/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b, c, x, y int
		_, _ = fmt.Scan(&a, &b, &c, &x, &y)
		x = max(x-a, 0)
		y = max(y-b, 0)
		if c >= x+y {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
