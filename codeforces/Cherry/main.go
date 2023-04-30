package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1554/A
func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n, prev, max int
		_, _ = fmt.Fscan(reader, &n)
		for j := 0; j < n; j++ {
			var a int
			_, _ = fmt.Fscan(reader, &a)
			exec := a * prev
			if max < exec {
				max = exec
			}
			prev = a
		}
		fmt.Println(max)
	}
}
