package main

import "fmt"

// https://codeforces.com/problemset/problem/510/A
func main() {
	var n, m int
	_, _ = fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		divisorOfFour := i%4 == 0
		divisorOfTwo := i%2 == 0
		if divisorOfFour && divisorOfTwo {
			fmt.Print("#")
		}
		for j := 1; j <= m; j++ {
			if (divisorOfTwo || divisorOfFour) && j < m {
				fmt.Print(".")
			} else if !divisorOfTwo {
				fmt.Print("#")
			}
		}
		if !divisorOfFour && divisorOfTwo {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
