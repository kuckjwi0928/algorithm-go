package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1810/A
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

		found := false
		l := len(arr)

		for i := 0; i < l; i++ {
			if i >= arr[i]-1 {
				found = true
				break
			}
		}

		if found {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
