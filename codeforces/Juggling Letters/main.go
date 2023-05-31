package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1397/A
func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	_, _ = fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var n int
		m := make(map[rune]int)
		_, _ = fmt.Fscan(reader, &n)
		for j := 0; j < n; j++ {
			var s string
			_, _ = fmt.Fscan(reader, &s)
			for _, c := range s {
				m[c]++
			}
		}

		yes := true

		for _, v := range m {
			if v%n != 0 {
				yes = false
				break
			}
		}

		if yes {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
