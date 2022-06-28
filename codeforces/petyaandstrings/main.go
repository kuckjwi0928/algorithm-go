package main

import (
	"fmt"
	"strings"
)

// https://codeforces.com/contest/112/problem/A
func main() {
	var l1, l2 string
	_, _ = fmt.Scan(&l1)
	_, _ = fmt.Scan(&l2)

	lower1, lower2 := strings.ToLower(l1), strings.ToLower(l2)

	if lower1 == lower2 {
		fmt.Println(0)
	} else if lower1 < lower2 {
		fmt.Println(-1)
	} else {
		fmt.Println(1)
	}
}
