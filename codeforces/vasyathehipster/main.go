package main

import "fmt"

// https://codeforces.com/problemset/problem/581/A
func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)

	count := 0
	limit := a + b

	for i := 1; i <= limit; i++ {
		if a == 0 || b == 0 {
			break
		}
		count++
		a--
		b--
	}

	fmt.Println(count, (a+b)/2)
}
