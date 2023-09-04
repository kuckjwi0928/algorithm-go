package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/686/A
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, x int
	_, _ = fmt.Fscan(reader, &n, &x)
	distressed := 0
	for i := 0; i < n; i++ {
		var op string
		var num int
		_, _ = fmt.Fscan(reader, &op, &num)
		if op == "+" {
			x += num
		} else {
			if x >= num {
				x -= num
			} else {
				distressed++
			}
		}
	}
	fmt.Println(x, distressed)
}
