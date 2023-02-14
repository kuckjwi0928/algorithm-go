package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1624/A
func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	_, _ = fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		min := 1_000_000_000
		max := 0
		for j := 0; j < n; j++ {
			var a int
			_, _ = fmt.Fscan(reader, &a)
			if a < min {
				min = a
			}
			if a > max {
				max = a
			}
		}
		fmt.Println(max - min)
	}
}
