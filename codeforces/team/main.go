package main

import "fmt"

// https://codeforces.com/problemset/problem/231/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	count := 0

	for n > 0 {
		var a, b, c int

		_, _ = fmt.Scan(&a, &b, &c)

		check := make(map[int]int)

		check[a]++
		check[b]++
		check[c]++

		if check[1] >= 2 {
			count++
		}
		n--
	}

	fmt.Println(count)
}
