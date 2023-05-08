package main

import (
	"fmt"
	"strings"
)

// https://codeforces.com/problemset/problem/1692/C
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		_, _ = fmt.Scan()
		arr := make([][]string, 8)
		for j := 0; j < 8; j++ {
			var s string
			_, _ = fmt.Scan(&s)
			arr[j] = strings.Split(s, "")
		}
		var row, col int
		for j := 1; j+1 < 8; j++ {
			for k := 1; k+1 < 8; k++ {
				if arr[j][k] == "." {
					continue
				}
				if arr[j+1][k+1] == "." {
					continue
				}
				if arr[j+1][k-1] == "." {
					continue
				}
				if arr[j-1][k+1] == "." {
					continue
				}
				if arr[j+1][k-1] == "." {
					continue
				}
				row, col = j, k
			}
		}
		fmt.Println(row+1, " ", col+1)
	}
}
