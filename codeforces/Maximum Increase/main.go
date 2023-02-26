package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/702/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	num := arr[0]
	count := 1
	result := count

	for i := 1; i < n; i++ {
		if num >= arr[i] {
			result = max(result, count)
			count = 1
		} else {
			count++
		}
		num = arr[i]
	}

	fmt.Println(max(result, count))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
