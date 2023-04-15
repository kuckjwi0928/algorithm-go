package main

import "fmt"

// https://codeforces.com/problemset/problem/540/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)
	var a, b string
	_, _ = fmt.Scan(&a, &b)
	answer := 0
	for i := 0; i < n; i++ {
		d1 := int(a[i] - '0')
		d2 := int(b[i] - '0')
		if d1 < d2 {
			d1, d2 = d2, d1
		}
		answer += min(d1-d2, 10-d1+d2)
	}
	fmt.Println(answer)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
