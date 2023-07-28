package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1850/C
func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int

	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		for j := 0; j < 8; j++ {
			var s string
			_, _ = fmt.Fscan(reader, &s)
			for k := 0; k < 8; k++ {
				if s[k] != '.' {
					fmt.Print(string(s[k]))
				}
			}
		}
		fmt.Println()
	}
}
