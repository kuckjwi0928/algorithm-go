package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/344/A
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	arr := make([]int, n)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &arr[i])
	}

	count := 1

	for i := 1; i < n; i++ {
		if arr[i] != arr[i-1] {
			count++
		}
	}

	fmt.Println(count)
}
