package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1371/A
func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	_, _ = fmt.Fscan(reader, &t)
	for t > 0 {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		result := n / 2
		if n%2 == 1 {
			result++
		}
		fmt.Println(result)
		t--
	}
}
