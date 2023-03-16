package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://codeforces.com/problemset/problem/1676/B
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
		sort.Ints(arr)
		min := arr[0]
		eat := 0
		for j := 1; j < n; j++ {
			eat += arr[j] - min
		}
		fmt.Println(eat)
	}
}
