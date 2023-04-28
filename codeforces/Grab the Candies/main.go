package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/1807/B
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

		mihai := 0
		bianca := 0

		for j := len(arr) - 1; j >= 0; j-- {
			if arr[j]%2 == 0 {
				mihai += arr[j]
			} else {
				bianca += arr[j]
			}
		}

		if mihai > bianca {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
