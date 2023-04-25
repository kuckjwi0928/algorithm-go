package main

import "fmt"

// https://codeforces.com/problemset/problem/1703/C
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
		for j := 0; j < n; j++ {
			var b int
			var s string
			_, _ = fmt.Scan(&b)
			_, _ = fmt.Scan(&s)
			for _, c := range s {
				if c == 'D' {
					arr[j] = (arr[j] + 1) % 10
				} else {
					arr[j] = (arr[j] - 1 + 10) % 10
				}
			}
		}
		for _, v := range arr {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
