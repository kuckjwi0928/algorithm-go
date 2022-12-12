package main

import "fmt"

// https://codeforces.com/problemset/problem/9/A
func main() {
	var Y, W int
	_, _ = fmt.Scan(&Y, &W)
	arr := []string{"1/1", "5/6", "2/3", "1/2", "1/3", "1/6"}
	fmt.Println(arr[max(Y, W)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
