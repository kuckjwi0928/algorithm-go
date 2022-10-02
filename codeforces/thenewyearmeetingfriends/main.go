package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/723/A
func main() {
	arr := make([]int, 3)
	for i := 0; i < len(arr); i++ {
		_, _ = fmt.Scan(&arr[i])
	}

	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	fmt.Println(abs(arr[0]-arr[1]) + abs(arr[2]-arr[1]))
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
