package main

import (
	"fmt"
)

// https://codeforces.com/problemset/problem/1783/A
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
		swap(arr, 0, 1)
		swap(arr, 0, n-1)

		if arr[0] == arr[1] {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
			for i := 0; i < n; i++ {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()
		}
	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
