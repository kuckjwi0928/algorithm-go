package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// https://codeforces.com/problemset/problem/1473/A
func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	_, _ = fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var n, d int
		_, _ = fmt.Fscan(reader, &n, &d)
		a := make([]int, n)
		for j := 0; j < n; j++ {
			_, _ = fmt.Fscan(reader, &a[j])
		}
		sort.Ints(a)
		ok := false
		for j := 1; j < n; j++ {
			if a[j]+a[j-1] <= d {
				ok = true
				break
			}
		}
		if a[n-1] <= d || ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
