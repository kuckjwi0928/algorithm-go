package main

import (
	"fmt"
	"unicode"
)

// https://codeforces.com/problemset/problem/1644/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for t > 0 {
		var s string
		_, _ = fmt.Scan(&s)

		r := []rune(s)

		keys := make(map[rune]bool)

		ok := true

		for i := 0; i < len(s); i++ {
			if unicode.IsLower(r[i]) {
				keys[r[i]] = true
			} else {
				_, ok = keys[r[i]+32]
				if !ok {
					break
				}
			}
		}

		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		t--
	}
}

func abs(a, b uint8) uint8 {
	if a > b {
		return a - b
	}
	return b - a
}
