package main

import "fmt"

// https://codeforces.com/problemset/problem/136/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		var j int
		_, _ = fmt.Scan(&j)
		arr[j-1] = i + 1
	}

	for i := 0; i < n; i++ {
		if i+1 == n {
			fmt.Print(arr[i])
		} else {
			fmt.Print(arr[i], " ")
		}
	}
}
