package main

import "fmt"

// https://codeforces.com/problemset/problem/1791/B
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		pos := []int{0, 0}
		candy := []int{1, 1}
		var n int
		var s string
		_, _ = fmt.Scan(&n)
		_, _ = fmt.Scan(&s)
		candyPosMoved := false
		for j := 0; j < n; j++ {
			if s[j] == 'U' {
				pos[0]++
			} else if s[j] == 'D' {
				pos[0]--
			} else if s[j] == 'L' {
				pos[1]--
			} else if s[j] == 'R' {
				pos[1]++
			}
			if pos[0] == candy[0] && pos[1] == candy[1] {
				candyPosMoved = true
				break
			}
		}
		if candyPosMoved {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
