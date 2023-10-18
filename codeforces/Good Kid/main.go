package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/1873/B
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			_, _ = fmt.Scan(&arr[i])
		}
		sort.Ints(arr)
		arr[0]++
		result := 1
		for i := 0; i < n; i++ {
			result *= arr[i]
		}
		fmt.Println(result)
	}
}
