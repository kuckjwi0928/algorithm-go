package main

import "fmt"

// https://codeforces.com/problemset/problem/265/A
func main() {
	var s, t string
	_, _ = fmt.Scan(&s, &t)

	j := 0

	for i := 0; i < len(t); i++ {
		if t[i] == s[j] {
			j++
		}
	}

	fmt.Println(j + 1)
}
