package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/1294/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for t > 0 {
		var a, b, c, n int
		_, _ = fmt.Scan(&a, &b, &c, &n)
		arr := []int{a, b, c}
		sort.Ints(arr)
		n -= arr[2] - arr[0] + arr[2] - arr[1]
		if n >= 0 {
			if n%3 == 0 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		} else {
			fmt.Println("NO")
		}
		t--
	}
}
