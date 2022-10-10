package main

import "fmt"

// https://codeforces.com/problemset/problem/472/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Println(n-4, 4)
	} else {
		fmt.Println(n-9, 9)
	}
}
