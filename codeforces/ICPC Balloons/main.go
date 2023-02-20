package main

import "fmt"

// https://codeforces.com/problemset/problem/1703/B
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		m := make(map[rune]bool)
		var n int
		_, _ = fmt.Scan(&n)
		var s string
		_, _ = fmt.Scan(&s)
		sum := 0
		for _, c := range s {
			_, ok := m[c]
			if !ok {
				sum += 2
				m[c] = true
			} else {
				sum += 1
			}
		}
		fmt.Println(sum)
	}
}
