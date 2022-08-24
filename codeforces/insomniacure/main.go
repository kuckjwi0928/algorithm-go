package main

import "fmt"

// https://codeforces.com/problemset/problem/148/A
func main() {
	var k, l, m, n, d int
	_, _ = fmt.Scan(&k, &l, &m, &n, &d)

	count := 0

	for i := 1; i <= d; i++ {
		if i%k == 0 || i%l == 0 || i%m == 0 || i%n == 0 {
			count++
		}
	}

	fmt.Println(count)
}
