package main

import "fmt"

// https://codeforces.com/problemset/problem/1692/B
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)

		m := make(map[int]int)
		for j := 0; j < n; j++ {
			var k int
			_, _ = fmt.Scan(&k)
			m[k]++
		}
		even := 0
		for _, v := range m {
			if v%2 == 0 {
				even++
			}
		}
		fmt.Println(len(m) - even)
	}
}
