package main

import "fmt"

// https://codeforces.com/problemset/problem/1829/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var s string
		_, _ = fmt.Scan(&s)
		diffString := "codeforces"
		count := 0
		for j := 0; j < 10; j++ {
			if s[j] != diffString[j] {
				count++
			}
		}
		fmt.Println(count)
	}
}
