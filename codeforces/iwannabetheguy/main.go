package main

import "fmt"

// https://codeforces.com/problemset/problem/469/A
func main() {
	var level int
	_, _ = fmt.Scan(&level)

	m := make(map[int]bool)

	for i := 0; i < 2; i++ {
		var p int
		_, _ = fmt.Scan(&p)
		for j := 0; j < p; j++ {
			var a int
			_, _ = fmt.Scan(&a)
			m[a] = true
		}
	}

	if len(m) == level {
		fmt.Println("I become the guy.")
	} else {
		fmt.Println("Oh, my keyboard!")
	}
}
