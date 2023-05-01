package main

import "fmt"

// https://codeforces.com/problemset/problem/214/A
func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)

	pair := 0

	if a == 1 && b == 1 {
		fmt.Println(2)
		return
	}

	for i := 0; i <= a/2; i++ {
		for j := 0; j <= b/2; j++ {
			if i*i+j == a && i+j*j == b {
				pair++
			}
		}
	}

	fmt.Println(pair)
}
