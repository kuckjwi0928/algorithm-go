package main

import "fmt"

// https://codeforces.com/problemset/problem/1311/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for t > 0 {
		var a, b int
		_, _ = fmt.Scan(&a, &b)
		if a == b {
			fmt.Println(0)
		}
		if a > b {
			if (a-b)%2 == 0 {
				fmt.Println(1)
			} else {
				fmt.Println(2)
			}
		} else if a < b {
			if (b-a)%2 == 0 {
				fmt.Println(2)
			} else {
				fmt.Println(1)
			}
		}
		t--
	}
}
