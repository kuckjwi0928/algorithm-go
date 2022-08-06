package main

import "fmt"

// https://codeforces.com/problemset/problem/271/A
func main() {
	var y int
	_, _ = fmt.Scan(&y)
	y += 1

	for true {
		m := make(map[int]bool)
		next := false
		year := y
		for year != 0 {
			digit := year % 10
			if m[digit] {
				next = true
				break
			}
			m[digit] = true
			year /= 10
		}
		if !next {
			break
		}
		y++
	}

	fmt.Println(y)
}
