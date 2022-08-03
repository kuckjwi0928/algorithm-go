package main

import "fmt"

// https://codeforces.com/problemset/problem/41/A
func main() {
	var s1, s2 string
	_, _ = fmt.Scan(&s1)
	_, _ = fmt.Scan(&s2)

	if len(s1) != len(s2) {
		fmt.Println("NO")
	} else {
		index := 0
		for i := len(s2) - 1; i >= 0; i-- {
			if s2[i] != s1[index] {
				fmt.Println("NO")
				return
			}
			index++
		}
		fmt.Println("YES")
	}
}
