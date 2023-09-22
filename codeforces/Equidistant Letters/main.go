package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/1626/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Scan(&s)
		r := []rune(s)
		sort.Slice(r, func(i int, j int) bool { return r[i] < r[j] })
		fmt.Println(string(r))
	}
}
