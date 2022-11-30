package main

import "fmt"

// https://codeforces.com/problemset/problem/1692/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)
	for n > 0 {
		var a, b, c, d int
		_, _ = fmt.Scan(&a, &b, &c, &d)
		count := 0
		for _, n := range []int{b, c, d} {
			if a < n {
				count++
			}
		}
		n--
		fmt.Println(count)
	}
}
