package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1506/A
func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	var n, m, x int64
	_, _ = fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		_, _ = fmt.Fscan(reader, &n, &m, &x)
		row := (x - 1) % n
		col := (x - 1) / n
		fmt.Println(row*m + col + 1)
	}
}
