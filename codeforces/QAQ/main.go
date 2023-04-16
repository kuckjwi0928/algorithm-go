package main

import "fmt"

// https://codeforces.com/problemset/problem/894/A
func main() {
	var s string
	_, _ = fmt.Scan(&s)
	count := 0
	l := len(s)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if s[i] == 'Q' && s[j] == 'A' && s[k] == 'Q' {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
