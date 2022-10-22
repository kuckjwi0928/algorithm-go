package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1374/A
func main() {
	var t int
	_, _ = fmt.Scan(&t)

	scanner := bufio.NewReader(os.Stdin)

	for i := 0; i < t; i++ {
		var x, y, n int
		_, _ = fmt.Fscan(scanner, &x, &y, &n)

		if n%x < y {
			fmt.Println(n - (n % x) - (x - y))
		} else {
			fmt.Println(n - (n % x) + y)
		}
	}
}
