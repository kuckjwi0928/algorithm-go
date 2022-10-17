package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/758/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&arr[i])
	}

	sort.Ints(arr)

	l := len(arr)
	max := arr[l-1]
	s := 0
	for i := 0; i < l-1; i++ {
		s += max - arr[i]
	}
	fmt.Println(s)
}
