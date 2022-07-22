package main

import (
	"fmt"
	"strings"
)

// https://codeforces.com/problemset/problem/59/A
func main() {
	var s string
	_, _ = fmt.Scan(&s)

	upper := 0
	lower := 0

	for _, w := range s {
		if w >= 65 && w <= 90 {
			upper++
		} else {
			lower++
		}
	}

	if upper > lower {
		fmt.Println(strings.ToUpper(s))
	} else {
		fmt.Println(strings.ToLower(s))
	}
}
