package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/1399/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for t > 0 {
		var a int
		_, _ = fmt.Scan(&a)
		arr := make([]int, a)

		for i := 0; i < a; i++ {
			_, _ = fmt.Scan(&arr[i])
		}

		sort.Ints(arr)

		yes := true
		for i := 1; i < len(arr); i++ {
			if abs(arr[i]-arr[i-1]) > 1 {
				yes = false
				break
			}
		}

		if yes {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		t--
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
