package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1669/B
func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	_, _ = fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		m := make(map[int]int)
		for j := 0; j < n; j++ {
			var a int
			_, _ = fmt.Fscan(reader, &a)
			m[a]++
		}

		answer := -1

		for num, count := range m {
			if count >= 3 {
				answer = num
				break
			}
		}

		fmt.Println(answer)
	}
}
