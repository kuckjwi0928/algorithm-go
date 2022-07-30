package main

import "fmt"

// https://codeforces.com/problemset/problem/266/B
func main() {
	var n, t int
	_, _ = fmt.Scan(&n, &t)
	var s string
	_, _ = fmt.Scan(&s)
	r := []rune(s)

	for t > 0 {
		for i := 1; i < n; i++ {
			if r[i] == 'G' && r[i-1] == 'B' {
				r[i] = 'B'
				r[i-1] = 'G'
				i++
			}
		}
		t--
	}

	fmt.Println(string(r))
}
