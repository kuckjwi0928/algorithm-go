package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1325/B
func main() {
	var t int
	_, _ = fmt.Scan(&t)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < t; i++ {
		var n int
		_, _ = fmt.Fscan(reader, &n)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			_, _ = fmt.Fscan(reader, &arr[j])
		}
		fmt.Println(unique(arr))
	}
}

func unique(arr []int) int {
	m := map[int]bool{}
	for _, v := range arr {
		m[v] = true
	}
	return len(m)
}
