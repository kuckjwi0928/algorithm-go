package main

import "fmt"

// https://www.acmicpc.net/problem/1193
func main() {
	var X int
	_, _ = fmt.Scan(&X)

	line := 0
	maxLine := 0

	denominator := 0
	numerator := 0

	for maxLine < X {
		line++
		maxLine += line
	}

	g := maxLine - X

	if line%2 == 0 {
		numerator = line - g
		denominator = g + 1
	} else {
		numerator = g + 1
		denominator = line - g
	}

	fmt.Printf("%d/%d", numerator, denominator)
}
