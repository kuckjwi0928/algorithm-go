package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1714/B
func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Fscan(reader, &n)

		arr := make([]int, n)
		m := make(map[int]int)

		for j := 0; j < n; j++ {
			_, _ = fmt.Fscan(reader, &arr[j])
		}

		count := 0

		for j := n - 1; j >= 0; j-- {
			if m[arr[j]] >= 1 {
				count = j + 1
				break
			}
			m[arr[j]]++
		}

		fmt.Println(count)
	}
}
