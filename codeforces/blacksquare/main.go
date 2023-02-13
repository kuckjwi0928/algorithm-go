package main

import "fmt"

// https://codeforces.com/problemset/problem/431/A
func main() {
	a := make([]int, 4)
	for i := 0; i < 4; i++ {
		_, _ = fmt.Scan(&a[i])
	}
	var s string
	_, _ = fmt.Scan(&s)
	sum := 0
	for _, w := range s {
		sum += a[w-'0'-1]
	}
	fmt.Println(sum)
}
