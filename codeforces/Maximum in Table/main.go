package main

import "fmt"

// https://codeforces.com/problemset/problem/509/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				arr[i][j] = 1
			} else {
				arr[i][j] = arr[i-1][j] + arr[i][j-1]
			}
		}
	}
	fmt.Println(arr[n-1][n-1])
}
