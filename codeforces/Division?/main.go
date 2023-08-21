package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1669/A
func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		rating := 1
		if n <= 1399 {
			rating = 4
		} else if n <= 1599 {
			rating = 3
		} else if n <= 1899 {
			rating = 2
		}
		fmt.Printf("Division %d\n", rating)
	}
}
