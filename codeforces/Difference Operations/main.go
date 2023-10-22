package main

import "fmt"

// https://codeforces.com/problemset/problem/1708/A
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
		possible := true
		for i := 1; i < n; i++ {
			if arr[i]%arr[0] != 0 {
				possible = false
			}
		}
		if possible {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
