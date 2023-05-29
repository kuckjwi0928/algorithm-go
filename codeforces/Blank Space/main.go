package main

import "fmt"

// https://codeforces.com/problemset/problem/1829/B
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

		count := 0
		result := 0

		for j := 0; j < n; j++ {
			if arr[j] == 0 {
				count++
			} else {
				result = max(result, count)
				count = 0
			}
		}

		fmt.Println(max(result, count))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
