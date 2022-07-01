package main

import (
	"fmt"
)

// https://codeforces.com/problemset/problem/281/A
func main() {
	var input string
	_, _ = fmt.Scan(&input)

	r := []rune(input)

	if r[0] >= 'a' && r[0] <= 'z' {
		r[0] = r[0] - 32
	}

	fmt.Println(string(r))
}
