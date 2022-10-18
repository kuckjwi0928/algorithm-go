package main

import "fmt"

// https://codeforces.com/problemset/problem/1512/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		var n int
		_, _ = fmt.Scan(&n)

		arr := make([]int, n)

		for j := 0; j < n; j++ {
			_, _ = fmt.Scan(&arr[j])
		}

		for j := 1; j < n; j++ {
			if j == 1 && arr[j] != arr[0] && arr[j] == arr[j+1] {
				fmt.Println(1)
				break
			}
			if arr[j] != arr[0] {
				fmt.Println(j + 1)
				break
			}
		}
	}
}
