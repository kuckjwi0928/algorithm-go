package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/1742/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for t > 0 {
		arr := make([]int, 3)
		for i := 0; i < len(arr); i++ {
			_, _ = fmt.Scan(&arr[i])
		}
		sort.Ints(arr)
		if arr[0]+arr[1] == arr[2] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		t--
	}
}
