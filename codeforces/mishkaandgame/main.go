package main

import "fmt"

// https://codeforces.com/problemset/problem/703/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	mishka := 0
	chris := 0

	for i := 0; i < n; i++ {
		var m, c int
		_, _ = fmt.Scan(&m, &c)
		if m > c {
			mishka++
		} else if c > m {
			chris++
		}
	}

	if mishka > chris {
		fmt.Println("Mishka")
	} else if chris > mishka {
		fmt.Println("Chris")
	} else {
		fmt.Println("Friendship is magic!^^")
	}
}
