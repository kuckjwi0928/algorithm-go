package main

import "fmt"

// https://codeforces.com/problemset/problem/155/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	arr := make([]int, n)
	m := make(map[int]bool)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&arr[i])
	}

	max := arr[0]
	min := arr[0]

	for i := 1; i < n; i++ {
		if max < arr[i] {
			m[arr[i]] = true
			max = arr[i]
		} else if min > arr[i] {
			m[arr[i]] = true
			min = arr[i]
		}
	}

	fmt.Println(len(m))
}
