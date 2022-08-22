package main

import "fmt"

// https://codeforces.com/problemset/problem/228/A
func main() {
	m := make(map[int]bool)

	for i := 0; i < 4; i++ {
		var j int
		_, _ = fmt.Scan(&j)
		m[j] = true
	}

	fmt.Println(4 - len(m))
}
