package main

import "fmt"

// https://codeforces.com/problemset/problem/1303/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Scan(&s)

		start := 0
		end := 0

		for i := 0; i < len(s); i++ {
			if s[i] == '1' {
				start = i
				break
			}
		}

		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '1' {
				end = i
				break
			}
		}

		count := 0

		for i := start; i < end; i++ {
			if s[i] == '0' {
				count++
			}
		}

		fmt.Println(count)
	}
}
