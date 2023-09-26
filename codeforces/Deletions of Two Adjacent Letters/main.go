package main

import "fmt"

// https://codeforces.com/problemset/problem/1650/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var s string
		var c string
		_, _ = fmt.Scan(&s, &c)
		found := false
		for j := 0; j < len(s); j++ {
			if string(s[j]) == c {
				left := j
				right := len(s) - j - 1
				if left%2 == 0 && right%2 == 0 {
					found = true
					break
				}
			}
		}
		if found {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
