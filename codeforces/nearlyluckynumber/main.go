package main

import "fmt"

// https://codeforces.com/problemset/problem/110/A
func main() {
	var n int64
	_, _ = fmt.Scan(&n)

	count := 0

	for n != 0 {
		m := n % 10
		if m == 4 || m == 7 {
			count++
		}
		n /= 10
	}

	if count == 4 || count == 7 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
