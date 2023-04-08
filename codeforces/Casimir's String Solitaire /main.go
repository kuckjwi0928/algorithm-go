package main

import "fmt"

// https://codeforces.com/problemset/problem/1579/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Scan(&s)

		m := make(map[rune]int)

		for _, w := range s {
			m[w]++
		}

		if m['B'] == m['A']+m['C'] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
