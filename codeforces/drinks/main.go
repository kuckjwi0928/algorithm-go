package main

import "fmt"

// https://codeforces.com/problemset/problem/200/B
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	sum := 0.0

	for i := 0; i < n; i++ {
		var p int
		_, _ = fmt.Scan(&p)
		sum += float64(p)
	}

	fmt.Printf("%.12f", sum/float64(n))
}
