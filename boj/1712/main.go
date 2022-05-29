package main

import "fmt"

// https://www.acmicpc.net/problem/1712
func main() {
	var A, B, C int
	_, _ = fmt.Scan(&A, &B, &C)
	if B >= C {
		fmt.Println(-1)
		return
	}
	fmt.Println(A/(C-B) + 1)
}
