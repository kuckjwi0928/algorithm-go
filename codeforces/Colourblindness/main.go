package main

import (
	"fmt"
	"strings"
)

// https://codeforces.com/problemset/problem/1722/B
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Scan(&n)
		arr := make([]string, 2)
		for j := 0; j < len(arr); j++ {
			var s string
			_, _ = fmt.Scan(&s)
			arr[j] = strings.ReplaceAll(s, "G", "B")
		}

		same := true

		for i := 0; i < len(arr[0]); i++ {
			if arr[0][i] != arr[1][i] {
				same = false
				break
			}
		}

		if same {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
