package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/1154/A
func main() {
	arr := make([]int, 4)

	for i := 0; i < len(arr); i++ {
		_, _ = fmt.Scan(&arr[i])
	}

	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	fmt.Println(arr[3]-arr[0], arr[3]-arr[1], arr[3]-arr[2])
}
