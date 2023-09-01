package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1833/A
func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	_, _ = fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var (
			n int
			s string
		)
		_, _ = fmt.Fscan(reader, &n, &s)
		m := make(map[string]int)
		for j := 0; j < n-1; j++ {
			m[s[j:j+2]]++
		}
		fmt.Println(len(m))
	}
}
