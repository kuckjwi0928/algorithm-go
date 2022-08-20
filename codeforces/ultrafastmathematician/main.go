package main

import "fmt"

// https://codeforces.com/problemset/problem/61/A
func main() {
	var a, b string
	_, _ = fmt.Scan(&a)
	_, _ = fmt.Scan(&b)

	r1 := []rune(a)
	r2 := []rune(b)

	ans := make([]rune, len(a))

	for i, w := range r1 {
		if w != r2[i] {
			ans[i] = '1'
		} else {
			ans[i] = '0'
		}
	}

	fmt.Println(string(ans))
}
