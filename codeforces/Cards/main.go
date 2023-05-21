package main

import "fmt"

// https://codeforces.com/problemset/problem/1220/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	var s string
	_, _ = fmt.Scan(&s)
	zero := 0
	one := 0
	for _, w := range s {
		if w == 'r' {
			zero++
		} else if w == 'n' {
			one++
		}
	}
	for i := 0; i < one; i++ {
		fmt.Print(1, " ")
	}
	for i := 0; i < zero; i++ {
		fmt.Print(0, " ")
	}
}
