package main

import "fmt"

// https://codeforces.com/problemset/problem/705/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print("I hate ")
		} else {
			fmt.Print("I love ")
		}
		if i != n-1 {
			fmt.Print("that ")
		} else {
			fmt.Print("it ")
		}
	}
}
