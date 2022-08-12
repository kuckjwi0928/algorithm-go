package main

import "fmt"

// https://codeforces.com/problemset/problem/1030/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)
	easyCount := 0
	hardCount := 0
	for i := 0; i < n; i++ {
		var j int
		_, _ = fmt.Scan(&j)
		if j == 0 {
			easyCount++
		} else {
			hardCount++
		}
	}
	if hardCount == easyCount {
		fmt.Println("EASY")
	} else if hardCount >= 1 {
		fmt.Println("HARD")
	} else {
		fmt.Println("EASY")
	}
}
