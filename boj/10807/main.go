package main

import "fmt"

// https://www.acmicpc.net/problem/10807
func main() {
	var n int
	_, _ = fmt.Scan(&n)

	counter := make(map[int]int)

	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Scan(&v)
		counter[v]++
	}

	var v int
	_, _ = fmt.Scan(&v)
	fmt.Println(counter[v])
}
