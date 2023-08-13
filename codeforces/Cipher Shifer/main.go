package main

import (
	"fmt"
	"strings"
)

// https://codeforces.com/problemset/problem/1840/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var (
			n int
			s string
		)

		_, _ = fmt.Scan(&n, &s)

		l := len(s)
		var sb strings.Builder

		j := 0

		for j < l {
			sb.WriteByte(s[j])
			k := j + 1
			for k < l && s[j] != s[k] {
				k++
			}
			j = k + 1
		}

		fmt.Println(sb.String())
	}
}
