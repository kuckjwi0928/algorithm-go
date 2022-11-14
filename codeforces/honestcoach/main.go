package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/1360/B
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for t > 0 {
		var n int
		_, _ = fmt.Scan(&n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			_, _ = fmt.Scan(&arr[i])
		}
		sort.Ints(arr)
		diff := 0
		dp := make([]int, n-1)
		for i := 1; i < len(arr); i++ {
			diff = arr[i] - arr[i-1]
			dp[i-1] = diff
		}
		sort.Ints(dp)
		fmt.Println(dp[0])
		t--
	}
}
