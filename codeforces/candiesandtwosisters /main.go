package main

import (
	"fmt"
	"math"
)

// https://codeforces.com/problemset/problem/1335/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n float64
		_, _ = fmt.Scan(&n)
		fmt.Println(int(math.Ceil(n/2)) - 1)
	}
}
