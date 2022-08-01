package main

import "fmt"

// https://codeforces.com/problemset/problem/734/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)
	var s string
	_, _ = fmt.Scan(&s)

	anton := 0
	danik := 0

	for _, r := range s {
		if r == 'A' {
			anton++
		} else if r == 'D' {
			danik++
		}
	}

	if anton == danik {
		fmt.Println("Friendship")
	} else if anton > danik {
		fmt.Println("Anton")
	} else {
		fmt.Println("Danik")
	}
}
