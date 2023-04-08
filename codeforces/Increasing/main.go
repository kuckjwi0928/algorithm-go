package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/1742/B
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)

		arr := make([]int, n)

		for j := 0; j < n; j++ {
			_, _ = fmt.Scan(&arr[j])
		}

		sort.Ints(arr)

		increasing := true

		for j := 0; j < n-1; j++ {
			if arr[j] >= arr[j+1] {
				increasing = false
				break
			}
		}

		if increasing {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
