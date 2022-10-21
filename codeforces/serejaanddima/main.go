package main

import (
	"fmt"
)

// https://codeforces.com/problemset/problem/381/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&arr[i])
	}

	sereja := 0
	dima := 0

	left := 0
	right := n - 1
	i := 1

	for left <= right {
		if i%2 == 0 {
			if arr[left] > arr[right] {
				dima += arr[left]
				left++
			} else {
				dima += arr[right]
				right--
			}
		} else {
			if arr[left] > arr[right] {
				sereja += arr[left]
				left++
			} else {
				sereja += arr[right]
				right--
			}
		}
		i++
	}

	fmt.Println(sereja, dima)
}
