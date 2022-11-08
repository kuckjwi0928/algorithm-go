package main

import (
	"fmt"
	"strings"
)

// https://codeforces.com/problemset/problem/1703/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for t > 0 {
		var s string
		_, _ = fmt.Scan(&s)
		if strings.ToLower(s) == "yes" {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		t--
	}
}
