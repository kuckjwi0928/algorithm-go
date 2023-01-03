package main

import "fmt"

// https://codeforces.com/problemset/problem/1520/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for t > 0 {
		suspicious := false

		var n int
		var s string
		_, _ = fmt.Scan(&n)
		_, _ = fmt.Scan(&s)

		m := make(map[uint8]int)
		m[s[0]] = 1

		for i := 1; i < n; i++ {
			if s[i] != s[i-1] && m[s[i]] > 0 {
				suspicious = true
				break
			}
			m[s[i]]++
		}

		if suspicious {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}

		t--
	}
}
