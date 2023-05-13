package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://codeforces.com/problemset/problem/1760/C
func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	_, _ = fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		arr := make([]int, n)
		sorted := make([]int, n)
		for j := 0; j < n; j++ {
			var a int
			_, _ = fmt.Fscan(reader, &a)
			arr[j] = a
			sorted[j] = a
		}
		sort.Ints(sorted)
		for j := 0; j < n; j++ {
			k := n - 1
			if arr[j] == sorted[k] {
				k = k - 1
			}
			fmt.Print(arr[j]-sorted[k], " ")
		}
		fmt.Println()
	}
}
