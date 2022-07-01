package main

import "fmt"

// https://codeforces.com/problemset/problem/236/A
func main() {
	var input string
	_, _ = fmt.Scan(&input)

	m := make(map[rune]bool)

	for _, in := range input {
		m[in] = true
	}

	if len(m)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
