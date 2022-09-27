package main

import "fmt"

// https://codeforces.com/problemset/problem/732/A
func main() {
	var k, r int
	_, _ = fmt.Scan(&k, &r)

	acc := k
	count := 1

	for acc%10 != r && acc%10 != 0 {
		acc += k
		count++
	}

	fmt.Println(count)
}
