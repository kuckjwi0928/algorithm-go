package main

import "fmt"

// https://codeforces.com/problemset/problem/1741/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var a, b string
		_, _ = fmt.Scan(&a, &b)
		if a == b {
			fmt.Println("=")
		} else if sizeCalculate(a) < sizeCalculate(b) {
			fmt.Println("<")
		} else {
			fmt.Println(">")
		}
	}
}

func sizeCalculate(s string) int {
	if s == "M" {
		return 0
	}
	size := 1
	for _, w := range s {
		if w == 'X' {
			size++
		}
	}
	if s[len(s)-1] == 'S' {
		return size * -1
	}
	return size
}
