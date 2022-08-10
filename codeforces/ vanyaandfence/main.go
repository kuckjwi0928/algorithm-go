package main

import "fmt"

// https://codeforces.com/problemset/problem/677/A
func main() {
	var n, h int
	_, _ = fmt.Scan(&n, &h)
	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		var input int
		_, _ = fmt.Scan(&input)
		arr = append(arr, input)
	}
	sum := 0
	for _, a := range arr {
		if h >= a {
			sum++
		} else {
			sum = sum + 2
		}
	}
	fmt.Println(sum)
}
