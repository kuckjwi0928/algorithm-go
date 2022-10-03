package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1409/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	reader := bufio.NewReader(os.Stdin)
	for t > 0 {
		var a, b int
		_, _ = fmt.Fscan(reader, &a, &b)
		if a < b {
			a, b = b, a
		}
		val := a - b
		result := val / 10
		if val%10 != 0 {
			result++
		}
		fmt.Println(result)
		t--
	}
}
