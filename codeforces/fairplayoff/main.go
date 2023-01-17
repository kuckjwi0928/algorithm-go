package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/1535/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for t > 0 {
		arr := make([]int, 4)

		for i := 0; i < len(arr); i++ {
			_, _ = fmt.Scan(&arr[i])
		}

		winner1 := max(arr[0], arr[1])
		winner2 := max(arr[2], arr[3])

		sort.Ints(arr)

		if (winner1 == arr[2] && winner2 == arr[3]) || (winner1 == arr[3] && winner2 == arr[2]) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		t--
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
