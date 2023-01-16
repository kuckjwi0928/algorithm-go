package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/1760/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for t > 0 {
		arr := make([]int, 3)
		_, _ = fmt.Scan(&arr[0], &arr[1], &arr[2])
		sort.Ints(arr)
		fmt.Println(arr[1])
		t--
	}
}
