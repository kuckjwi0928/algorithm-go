package main

import "fmt"

// https://codeforces.com/problemset/problem/1551/B1
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Scan(&s)
		m := make(map[rune]int)
		answer := 0
		for _, w := range s {
			m[w]++
		}
		for k, v := range m {
			if v >= 2 {
				answer++
				delete(m, k)
			}
		}
		fmt.Println(len(m)/2 + answer)
	}
}
