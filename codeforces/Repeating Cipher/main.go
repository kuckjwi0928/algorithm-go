package main

import "fmt"

// https://codeforces.com/problemset/problem/1095/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	var s string
	_, _ = fmt.Scan(&s)

	decryptString := string(s[0])

	start := 2
	for i := start; i <= n; i += start {
		decryptString += string(s[i-1])
		start++
	}

	fmt.Println(decryptString)
}
